package entity

type ServiceOutputInterface interface {
     SaveServiceOutput(serviceOutput *ServiceOutput, service string) error
     FindOneByIdAndService(id string, service string) (*ServiceOutput, error)
     FindAllByService(service string) ([]*ServiceOutput, error)
     FindAllByServiceAndSource(service string, source string) ([]*ServiceOutput, error)
}
