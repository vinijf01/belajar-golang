package simple

type SimpleRepository struct {
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpeRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{SimpleRepository: repository}
}
