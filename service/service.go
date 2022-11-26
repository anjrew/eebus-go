package service

import (
	"crypto/x509"
	"fmt"
	"sync"

	"github.com/DerAndereAndi/eebus-go/logging"
	"github.com/DerAndereAndi/eebus-go/spine"
	"github.com/DerAndereAndi/eebus-go/spine/model"
)

type EEBUSServiceDelegate interface {
	// RemoteServicesListUpdated(services []ServiceDetails)

	// report a connection to a SKI
	RemoteSKIConnected(service *EEBUSService, ski string)

	// report a disconnection to a SKI
	RemoteSKIDisconnected(service *EEBUSService, ski string)
}

// A service is the central element of an EEBUS service
// including its websocket server and a zeroconf service.
type EEBUSService struct {
	ServiceDescription *ServiceDescription

	// The local service details
	LocalService *ServiceDetails

	// Connection Registrations
	connectionsHub *connectionsHub

	// The SPINE specific device definition
	spineLocalDevice *spine.DeviceLocalImpl

	serviceDelegate EEBUSServiceDelegate

	startOnce sync.Once
}

// creates a new EEBUS service
func NewEEBUSService(ServiceDescription *ServiceDescription, serviceDelegate EEBUSServiceDelegate) *EEBUSService {
	return &EEBUSService{
		ServiceDescription: ServiceDescription,
		serviceDelegate:    serviceDelegate,
	}
}

// Sets a custom logging implementation
// By default NoLogging is used, so no logs are printed
func (s *EEBUSService) SetLogging(logger logging.Logging) {
	if logger == nil {
		return
	}
	logging.Log = logger
}

// Starts the service by initializeing mDNS and the server.
func (s *EEBUSService) Setup() error {
	if s.ServiceDescription.port == 0 {
		s.ServiceDescription.port = defaultPort
	}

	sd := s.ServiceDescription

	leaf, err := x509.ParseCertificate(sd.certificate.Certificate[0])
	if err != nil {
		logging.Log.Error(err)
		return err
	}

	ski, err := skiFromCertificate(leaf)
	if err != nil {
		logging.Log.Error(err)
		return err
	}

	s.LocalService = &ServiceDetails{
		SKI:                ski,
		ShipID:             sd.shipIdentifier(),
		deviceType:         sd.deviceType,
		registerAutoAccept: sd.registerAutoAccept,
	}

	logging.Log.Info("Local SKI: ", ski)

	vendor := sd.vendorCode
	if vendor == "" {
		vendor = sd.deviceBrand
	}

	serial := sd.deviceSerialNumber
	if serial != "" {
		serial = fmt.Sprintf("-%s", serial)
	}

	// Create the SPINE device address, according to Protocol Specification 7.1.1.2
	deviceAdress := fmt.Sprintf("d:_i:%s_%s%s", vendor, sd.deviceModel, serial)

	// Create the local SPINE device
	s.spineLocalDevice = spine.NewDeviceLocalImpl(
		sd.deviceBrand,
		sd.deviceModel,
		sd.deviceSerialNumber,
		sd.shipIdentifier(),
		deviceAdress,
		sd.deviceType,
		sd.featureSet,
	)

	// Create the device entity and add it to the SPINE device
	entityAddress := []model.AddressEntityType{1}
	var entityType model.EntityTypeType
	switch sd.deviceType {
	case model.DeviceTypeTypeEnergyManagementSystem:
		entityType = model.EntityTypeTypeCEM
	case model.DeviceTypeTypeChargingStation:
		entityType = model.EntityTypeTypeEVSE
	default:
		logging.Log.Errorf("Unknown device type: %s", sd.deviceType)
	}
	entity := spine.NewEntityLocalImpl(s.spineLocalDevice, entityType, entityAddress)
	s.spineLocalDevice.AddEntity(entity)

	// Setup connections hub with mDNS and websocket connection handling
	hub, err := newConnectionsHub(s.spineLocalDevice, s.ServiceDescription, s.LocalService)
	if err != nil {
		return err
	}

	s.connectionsHub = hub

	return nil
}

// Starts the service
func (s *EEBUSService) Start() {
	s.startOnce.Do(func() {
		s.connectionsHub.start()
	})
}

// Shutdown all services and stop the server.
func (s *EEBUSService) Shutdown() {
	// Shut down all running connections
	s.connectionsHub.shutdown()
}

func (s *EEBUSService) LocalDevice() *spine.DeviceLocalImpl {
	return s.spineLocalDevice
}

// return the local entity 1
func (s *EEBUSService) LocalEntity() *spine.EntityLocalImpl {
	return s.spineLocalDevice.Entity([]model.AddressEntityType{1})
}

// Add a new entity, used for connected EVs
// Only for EVSE implementations
func (s *EEBUSService) AddEntity(entity *spine.EntityLocalImpl) {
	s.spineLocalDevice.AddEntity(entity)
}

// Remove an entity, used for disconnected EVs
// Only for EVSE implementations
func (s *EEBUSService) RemoveEntity(entity *spine.EntityLocalImpl) {
	s.spineLocalDevice.RemoveEntity(entity)
}

// return all remote devices
func (s *EEBUSService) RemoteDevices() []*spine.DeviceRemoteImpl {
	return s.spineLocalDevice.RemoteDevices()
}

func (s *EEBUSService) RemoteDeviceForSki(ski string) *spine.DeviceRemoteImpl {
	return s.spineLocalDevice.RemoteDeviceForSki(ski)
}

// return a specific remote device of a given DeviceType
func (s *EEBUSService) RemoteDeviceOfType(deviceType model.DeviceTypeType) *spine.DeviceRemoteImpl {
	for _, device := range s.spineLocalDevice.RemoteDevices() {
		if *device.DeviceType() == deviceType {
			return device
		}
	}
	return nil
}

// Adds a new device to the list of known devices which can be connected to
// and connect it if it is currently not connected
func (s *EEBUSService) PairRemoteService(service ServiceDetails) {
	s.connectionsHub.PairRemoteService(service)
}

// Returns if the provided SKI is from a registered service
func (s *EEBUSService) IsRemoteServiceForSKIPaired(ski string) bool {
	return s.connectionsHub.IsRemoteServiceForSKIPaired(ski)
}

// Remove a device from the list of known devices which can be connected to
// and disconnect it if it is currently connected
func (s *EEBUSService) UnpairRemoteService(ski string) error {
	return s.connectionsHub.UnpairRemoteService(ski)
}

// Close a connection to a remote SKI
func (s *EEBUSService) DisconnectSKI(ski string) {
	s.connectionsHub.DisconnectSKI(ski)
}
