package api

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

//Leases can be used to query the Leases routes
type Leases struct {
	c *Client
}

//Leases returns a handle to the Leases routes
func (c *Client) Leases() *Leases {
	return &Leases{c}
}

// Index returns a list of Leases.
func (s *Leases) Index(rID string) (resources.LeasesV1, error) {
	rLeases, err := s.c.ReceiveResource("GET", "/reservations"+rID+"/leases", "", "")
	if err != nil {
		return resources.LeasesV1{}, err
	}
	if leases, ok := rLeases.(*resources.LeasesV1); ok {
		return *leases, nil
	}
	return resources.LeasesV1{}, errors.New("Lease Index call error.")
}

// Show returns the requested Lease.
func (s *Leases) Show(lID string, rLeaseIn resources.LeaseV1) (resources.LeaseV1, error) {
	rLeaseOut, err := s.c.ReceiveResource("GET", "/leases/"+lID, rLeaseIn.Type(), rLeaseIn.Version())
	if err != nil {
		return resources.LeaseV1{}, err
	}
	if lease, ok := rLeaseOut.(*resources.LeaseV1); ok {
		return *lease, nil
	}
	return resources.LeaseV1{}, errors.New("Lease Show call error.")
}

// Update updates the requested Lease and returns its location.
func (s *Leases) Update(lID string, rLease resources.LeaseV1) (string, error) {
	lLocation, err := s.c.SendResource("PATCH", "/leases/"+lID, &rLease)
	if err != nil {
		return "", err
	}
	return lLocation, nil
}

// UpdateShowLease updates a Lease and then returns that Lease.
func (s *Leases) UpdateShowLease(lID string, rLease resources.LeaseV1) (resources.LeaseV1, error) {
	rLeaseOut, err := s.c.SendReceiveResource("PATCH", "GET", "/leases/"+lID, &rLease)
	if err != nil {
		return resources.LeaseV1{}, err
	}
	if lease, ok := rLeaseOut.(*resources.LeaseV1); ok {
		return *lease, nil
	}
	return resources.LeaseV1{}, errors.New("UpdateShowLease call error.")
}