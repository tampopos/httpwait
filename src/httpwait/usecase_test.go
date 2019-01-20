package httpwait

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tampopos/httpwait/src/client"
)

func TestOnFirstSuccessByStatusCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	args := &WaitArgs{
		Request: client.Request{
			Method:  "GET",
			Timeout: 60,
			URL:     "http://example.com/",
		},
		StatusCode: 200,
	}

	mockStopwatch := NewMockStopwatch(ctrl)
	mockStopwatch.EXPECT().Start().Return(mockStopwatch)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().GetStatusCode(gomock.Eq(&args.Request)).Return(200, nil)

	useCase := CreateUseCase(mockStopwatch, mockClient)
	err := useCase.Wait(args)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
func TestOnSecondSuccessByStatusCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	args := &WaitArgs{
		Request: client.Request{
			Method:  "GET",
			Timeout: 60,
			URL:     "http://example.com/",
		},
		StatusCode: 200,
	}

	mockStopwatch := NewMockStopwatch(ctrl)
	mockStopwatch.EXPECT().Start().Return(mockStopwatch)
	mockStopwatch.EXPECT().GetElapsedSeconds().Return(float64(10))

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().GetStatusCode(gomock.Eq(&args.Request)).Return(404, nil).Times(1)
	mockClient.EXPECT().GetStatusCode(gomock.Eq(&args.Request)).Return(200, nil).Times(1)

	useCase := CreateUseCase(mockStopwatch, mockClient)
	err := useCase.Wait(args)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
func TestOnFirstSuccessByBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	args := &WaitArgs{
		Request: client.Request{
			Method:  "GET",
			Timeout: 60,
			URL:     "http://example.com/",
		},
		Result:     "200",
		StatusCode: -1,
	}

	mockStopwatch := NewMockStopwatch(ctrl)
	mockStopwatch.EXPECT().Start().Return(mockStopwatch)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().GetBody(gomock.Eq(&args.Request)).Return("200", nil)

	useCase := CreateUseCase(mockStopwatch, mockClient)
	err := useCase.Wait(args)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
func TestOnSecondSuccessByBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	args := &WaitArgs{
		Request: client.Request{
			Method:  "GET",
			Timeout: 60,
			URL:     "http://example.com/",
		},
		Result:     "200",
		StatusCode: -1,
	}

	mockStopwatch := NewMockStopwatch(ctrl)
	mockStopwatch.EXPECT().Start().Return(mockStopwatch)
	mockStopwatch.EXPECT().GetElapsedSeconds().Return(float64(10))

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().GetBody(gomock.Eq(&args.Request)).Return("404", nil).Times(1)
	mockClient.EXPECT().GetBody(gomock.Eq(&args.Request)).Return("200", nil).Times(1)

	useCase := CreateUseCase(mockStopwatch, mockClient)
	err := useCase.Wait(args)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
func TestOnFirstTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	args := &WaitArgs{
		Request: client.Request{
			Method:  "GET",
			Timeout: 10,
			URL:     "http://example.com/",
		},
		StatusCode: 200,
	}

	mockStopwatch := NewMockStopwatch(ctrl)
	mockStopwatch.EXPECT().Start().Return(mockStopwatch)
	mockStopwatch.EXPECT().GetElapsedSeconds().Return(float64(10))

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().GetStatusCode(gomock.Eq(&args.Request)).Return(404, nil).AnyTimes()

	useCase := CreateUseCase(mockStopwatch, mockClient)
	err := useCase.Wait(args)

	if err.Error() != fmt.Errorf("Timeout").Error() {
		if err != nil {
			t.Fatalf("failed test %#v", err)
			return
		}
		t.Fatalf("failed test")
	}
}
func TestOnSecondTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	args := &WaitArgs{
		Request: client.Request{
			Method:  "GET",
			Timeout: 20,
			URL:     "http://example.com/",
		},
		StatusCode: 200,
	}

	mockStopwatch := NewMockStopwatch(ctrl)
	mockStopwatch.EXPECT().Start().Return(mockStopwatch)
	mockStopwatch.EXPECT().GetElapsedSeconds().Return(float64(10)).Times(1)
	mockStopwatch.EXPECT().GetElapsedSeconds().Return(float64(20)).Times(1)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().GetStatusCode(gomock.Eq(&args.Request)).Return(404, nil).AnyTimes()

	useCase := CreateUseCase(mockStopwatch, mockClient)
	err := useCase.Wait(args)
	if err.Error() != fmt.Errorf("Timeout").Error() {
		if err != nil {
			t.Fatalf("failed test %#v", err)
			return
		}
		t.Fatalf("failed test")
	}
}
func TestOnErrorByStatusCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	args := &WaitArgs{
		Request: client.Request{
			Method:  "GET",
			Timeout: 20,
			URL:     "http://example.com/",
		},
		StatusCode: 200,
	}

	mockStopwatch := NewMockStopwatch(ctrl)
	mockStopwatch.EXPECT().Start().Return(mockStopwatch)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().GetStatusCode(gomock.Eq(&args.Request)).Return(-1, fmt.Errorf("Error")).AnyTimes()

	useCase := CreateUseCase(mockStopwatch, mockClient)
	err := useCase.Wait(args)
	if err.Error() != fmt.Errorf("Error").Error() {
		if err != nil {
			t.Fatalf("failed test %#v", err)
			return
		}
		t.Fatalf("failed test")
	}
}
func TestOnErrorByBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	args := &WaitArgs{
		Request: client.Request{
			Method:  "GET",
			Timeout: 20,
			URL:     "http://example.com/",
		},
		Result:     "200",
		StatusCode: -1,
	}

	mockStopwatch := NewMockStopwatch(ctrl)
	mockStopwatch.EXPECT().Start().Return(mockStopwatch)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().GetBody(gomock.Eq(&args.Request)).Return("", fmt.Errorf("Error")).AnyTimes()

	useCase := CreateUseCase(mockStopwatch, mockClient)
	err := useCase.Wait(args)
	if err.Error() != fmt.Errorf("Error").Error() {
		if err != nil {
			t.Fatalf("failed test %#v", err)
			return
		}
		t.Fatalf("failed test")
	}
}
