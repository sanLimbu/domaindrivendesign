package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/sanLimbu/ddd-go/aggregate"
	"github.com/sanLimbu/ddd-go/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (m *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range m.products {
		products = append(products, product)
	}
	return products, nil
}

func (m *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := m.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (m *MemoryProductRepository) Add(newprod aggregate.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	m.products[newprod.GetID()] = newprod
	return nil
}

func (m *MemoryProductRepository) Update(upprod aggregate.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[upprod.GetID()]; ok {
		return product.ErrProductNotFound
	}

	m.products[upprod.GetID()] = upprod

	return nil
}

func (m *MemoryProductRepository) Delete(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[id]; ok {
		return product.ErrProductNotFound
	}

	delete(m.products, id)
	return nil

}
