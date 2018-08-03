package storage

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"twreporter.org/go-api/models"
)

// CreateAPayByPrimeDonation creates a draft order in database
func (g *GormStorage) CreateAPayByPrimeDonation(m models.PayByPrimeDonation) error {
	errWhere := "GormStorage.CreateAPayByPrimeDonation"
	err := g.db.Create(&m).Error
	if nil != err {
		log.Error(err.Error())
		return g.NewStorageError(err, errWhere, fmt.Sprintf("can not create the record(%#v)", m))
	}
	return nil
}

// UpdateAPayByPrimeDonation updates the draft record with the Tappay response data by order
func (g *GormStorage) UpdateAPayByPrimeDonation(order string, m models.PayByPrimeDonation) error {
	errWhere := "GormStorage.UpdateAPayByPrimeDonation"
	err := g.db.Model(&m).Where("order_number = ?", order).Updates(m).Error
	if nil != err {
		log.Error(err.Error())
		return g.NewStorageError(err, errWhere, fmt.Sprintf("can not update prime donation(order: %s, record: %#v)", order, m))
	}
	return nil
}

//TODO
func (g *GormStorage) CreateAPeriodDonation(m models.PeriodicDonation) error {
	return nil
}

//TODO
func (g *GormStorage) CreateAPayByCardTokenDonation(m models.PayByCardTokenDonation) error {
	return nil
}

//TODO
func (g *GormStorage) CreateAPayByOtherMethodDonation(m models.PayByOtherMethodDonation) error {
	return nil
}

//TODO
func (g *GormStorage) GetDonationsByPayMethods(filters []string, offset uint, limit uint) (models.DonationRecord, error) {
	return models.DonationRecord{}, nil
}