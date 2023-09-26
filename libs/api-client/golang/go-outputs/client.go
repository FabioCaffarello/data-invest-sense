package gooutputs

import (
     "context"
     "fmt"
     "net/http"

     inputDTO "libs/dtos/golang/dto-lake-outputs/input"
     outputDTO "libs/dtos/golang/dto-lake-outputs/output"
     gorequest "libs/golang/go-request"
)

type Client struct {
     ctx     context.Context
     baseURL string
}

func NewClient() *Client {
     return &Client{
          ctx:     context.Background(),
          baseURL: "http://lake-outputs:8000",
     }
}

func (c *Client) CreateOutput(service string, serviceOutput inputDTO.ServiceOutputDTO) (outputDTO.ServiceOutputDTO, error) {
     url := fmt.Sprintf("%s/outputs/service/%s", c.baseURL, service)
     req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, serviceOutput)
     if err != nil {
          return outputDTO.ServiceOutputDTO{}, err
     }

     var output outputDTO.ServiceOutputDTO
     if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
          return outputDTO.ServiceOutputDTO{}, err
     }

     return output, nil
}

