package customerrors

import "fmt"

type RepoError struct {
	Operation string
	Err       error
}

func (r *RepoError) Error()string {
	return fmt.Sprintf("repo error during %s : %v ",r.Operation,r.Err)
}

func NewRepoError(Operation string,Err error)*RepoError{
   return &RepoError{
	Operation: Operation,
	Err: Err,
   }
}