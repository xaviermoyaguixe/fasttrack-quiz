package storage

import "fasttrackquiz/types"

type Storage interface {
	Get() ([]types.Question, error)
	Submit(*types.SubmitRequest) error
}
