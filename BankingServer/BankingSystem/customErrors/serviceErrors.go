package customerrors

import "fmt"

type ServiceError struct {
	Operation string
	Err       error
}

func (r *ServiceError) Error()string {
	return fmt.Sprintf("service error during %s : %v ",r.Operation,r.Err)
}

func NewServiceError(Operation string,Err error)*ServiceError{
   return &ServiceError{
	Operation: Operation,
	Err: Err,
   }
}