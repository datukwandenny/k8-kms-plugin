/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	kmsapi "k8s.io/kms/apis/v2" // The Kubernetes KMS API
)

// GRPCService is a grpc server that runs the kms v2 alpha1 API.
type GRPCService struct {
	addr    string
	timeout time.Duration
	server  *grpc.Server
curl --request POST \
--url https://app.infisical.com/api/v1/kms/keys/<key-id>/encrypt \
--header 'Content-Type: application/json' \
--data '{
    "plaintext": "lUFHM5Ggwo6TOfpuN1S==" // base64 encoded plaintext
}'
	kmsService Service
}

var _ kmsapi.KeyManagementServiceServer = (*GRPCService)(nil)

// NewGRPCService creates an instance of GRPCService.
func NewGRPCService(
	address string,
	timeout time.Duration,
	kmsService Service,
) *GRPCService {
	return &GRPCService{
		addr:       address,
		timeout:    timeout,
		kmsService: kmsService,
	}
}

// ListenAndServe accepts incoming connections on a Unix socket. It is a blocking method.
// Returns non-nil error unless Close or Shutdown is called.
func (s *GRPCService) ListenAndServe() error {
	if !strings.HasPrefix(s.addr, "@") {
		err := os.Remove(s.addr)
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to delete the socket file, error: %w", err)
		}
	}

	ln, err := net.Listen("unix", s.addr)

	if err != nil {
		return err
	}
	defer ln.Close()

	gs := grpc.NewServer(
		grpc.ConnectionTimeout(s.timeout),
	)
	s.server = gs

	kmsapi.RegisterKeyManagementServiceServer(gs, s)

	log.Printf("Listening on socket: %s\n", s.addr)

	return gs.Serve(ln)
}

func (s *GRPCService) Start() chan error {
	errorCh := make(chan error)
	log.Println("Starting up plugin server...")
	go func() {
		defer func() {
			s.Shutdown()
			close(errorCh)
		}()
		select {
		case errorCh <- s.ListenAndServe():
		default:
		}
	}()

	return errorCh
}

// Shutdown performs a graceful shutdown. Doesn't accept new connections and
// blocks until all pending RPCs are finished.
func (s *GRPCService) Shutdown() {
	if s.server != nil {
		s.server.GracefulStop()
	}
}

// Close stops the server by closing all connections immediately and cancels
// all active RPCs.
func (s *GRPCService) Close() {
	if s.server != nil {
		s.server.Stop()
	}
}

// Status sends a status request to specified kms service.
func (s *GRPCService) Status(ctx context.Context, _ *kmsapi.StatusRequest) (*kmsapi.StatusResponse, error) {
	res, err := s.kmsService.Status(ctx)
	if err != nil {
		return nil, err
	{
  "success": true,
  "data": [
    {
      "block_id": 444832513,
      "trans_id": "2Vtu29XM5ki6wW7MSE4VTpa7C5ugnkGzVR3GwqPoCupDs5aZWXtDopUwBirh59dRhsuNpU91AJeVCcGK8YgDbMRV",
      "block_time": 1788710313,
      "activity_type": "ACTIVITY_SPL_TRANSFER",
      "from_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "from_token_account": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "to_address": "4Ft3mHekEXqpE9RtiFzU78RNdQnQSEHeVbGT3JQSYCk5",
      "to_token_account": "4Ft3mHekEXqpE9RtiFzU78RNdQnQSEHeVbGT3JQSYCk5",
      "token_address": "So11111111111111111111111111111111111111111",
      "token_decimals": 9,
      "amount": 189033280,
      "flow": "out",
      "value": 20.000809355377086,
      "source": [
        "11111111111111111111111111111111"
      ],
      "time": "2026-09-06T15:58:33.000Z"
    },
    {
      "block_id": 444826343,
      "trans_id": "3dJC1rpMcn4jBxVcYjgU8fLyVZTsrN5TT8DBVhaeUDPmpnFqfTSnUgdT7jjo6ys5jxAe1mDDEWTPZSte5CZKyMCZ",
      "block_time": 1788708354,
      "activity_type": "ACTIVITY_SPL_TRANSFER",
      "from_address": "6tckHFBpiJ8YgYN8FUskvtvTpXQZ55g5LHeo1kvELoDQ",
      "from_token_account": "6tckHFBpiJ8YgYN8FUskvtvTpXQZ55g5LHeo1kvELoDQ",
      "to_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "to_token_account": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "token_address": "So11111111111111111111111111111111111111111",
      "token_decimals": 9,
      "amount": 547490474,
      "flow": "in",
      "value": 57.73372700650698,
      "source": [
        "dahPEoZGXfyV58JqqH85okdHmpN8U2q8owgPUXSCPxe"
      ],
      "time": "2026-09-06T15:25:54.000Z"
    },
    {
      "block_id": 444822965,
      "trans_id": "33jy66iVzMDschRvnD6gZ6DNVD9nizeQukDX798qboaVm8Lo2Eyrh9v4zVXhtcMef7CTsykvF5rBrcX3veEB4NK",
      "block_time": 1788707282,
      "activity_type": "ACTIVITY_SPL_TRANSFER",
      "from_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "from_token_account": "CfXuMvpEn5U6494o8thkw8KrS6EXowwgxz8vS5mw8jqq",
      "to_address": "ZfctMHBkZNTqeYGE47ekxtydgXgpo9xKJCAasjaCLTU",
      "to_token_account": "7842wRnqHvLoNKqc5UarfXx6StngtFS95311CW8bEk5E",
      "token_address": "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB",
      "token_decimals": 6,
      "amount": 39405237,
      "flow": "out",
      "value": 39.405237,
      "source": [
        "FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"
      ],
      "time": "2026-09-06T15:08:02.000Z"
    },
    {
      "block_id": 444822682,
      "trans_id": "24KQmYLaabLXZPnPFXQjeqzpLvDBgCUs4BbevLcatn3NShitd2Z6LfvDVbnfWadmjJ2MJejvSUBVnEniggm1mHhY",
      "block_time": 1788707193,
      "activity_type": "ACTIVITY_SPL_CLOSE_ACCOUNT",
      "from_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "from_token_account": "3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA",
      "to_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "to_token_account": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "token_address": "So11111111111111111111111111111111111111111",
      "token_decimals": 9,
      "amount": 1855569,
      "flow": "out",
      "value": 0.19563743495415228,
      "source": [
        "FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"
      ],
      "time": "2026-09-06T15:06:33.000Z"
    },
    {
      "block_id": 444822682,
      "trans_id": "24KQmYLaabLXZPnPFXQjeqzpLvDBgCUs4BbevLcatn3NShitd2Z6LfvDVbnfWadmjJ2MJejvSUBVnEniggm1mHhY",
      "block_time": 1788707193,
      "activity_type": "ACTIVITY_SPL_TRANSFER",
      "from_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "from_token_account": "3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA",
      "to_address": "ZfctMHBkZNTqeYGE47ekxtydgXgpo9xKJCAasjaCLTU",
      "to_token_account": "8b9c7XAStj2dKpqk3U2SsXcxh5jPkuFTExUhnFZo4iFZ",
      "token_address": "So11111111111111111111111111111111111111112",
      "token_decimals": 9,
      "amount": 187573710,
      "flow": "out",
      "value": 19.776380985689038,
      "source": [
        "FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"
      ],
      "time": "2026-09-06T15:06:33.000Z"
    },
    {
      "block_id": 444822682,
      "trans_id": "24KQmYLaabLXZPnPFXQjeqzpLvDBgCUs4BbevLcatn3NShitd2Z6LfvDVbnfWadmjJ2MJejvSUBVnEniggm1mHhY",
      "block_time": 1788707193,
      "activity_type": "ACTIVITY_SPL_TRANSFER",
      "from_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "from_token_account": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "to_address": "3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA",
      "to_token_account": "3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA",
      "token_address": "So11111111111111111111111111111111111111111",
      "token_decimals": 9,
      "amount": 187573710,
      "flow": "out",
      "value": 19.776380985689038,
      "source": [
        "FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"
      ],
      "time": "2026-09-06T15:06:33.000Z"
    },
    {
      "block_id": 444822682,
      "trans_id": "24KQmYLaabLXZPnPFXQjeqzpLvDBgCUs4BbevLcatn3NShitd2Z6LfvDVbnfWadmjJ2MJejvSUBVnEniggm1mHhY",
      "block_time": 1788707193,
      "activity_type": "ACTIVITY_SPL_CREATE_ACCOUNT",
      "from_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "from_token_account": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "to_address": "3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA",
      "to_token_account": "3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA",
      "token_address": "So11111111111111111111111111111111111111111",
      "token_decimals": 9,
      "amount": 1855569,
      "flow": "out",
      "value": 0.19563743495415228,
      "source": [
        "FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"
      ],
      "time": "2026-09-06T15:06:33.000Z"
    },
    {
      "block_id": 444822405,
      "trans_id": "5MkB1gTizAtkbFrHyEGLnBpFNPnmMMThhkxL9fKWhaSLZ2P9aNJqNQadQtBDVXYgVvZVpPpJZ2Zm6bwyBmKEFgFS",
      "block_time": 1788707104,
      "activity_type": "ACTIVITY_SPL_TRANSFER",
      "from_address": "HMueTSsKEjyXEQAzyDMdx8mDPAUEaVA64gTaVJX2sfpP",
      "from_token_account": "HMueTSsKEjyXEQAzyDMdx8mDPAUEaVA64gTaVJX2sfpP",
      "to_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "to_token_account": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "token_address": "So11111111111111111111111111111111111111111",
      "token_decimals": 9,
      "amount": 1,
      "flow": "in",
      "value": 1.0529833027702624e-7,
      "source": [
        "11111111111111111111111111111111"
      ],
      "time": "2026-09-06T15:05:04.000Z"
    },
    {
      "block_id": 444822357,
      "trans_id": "34yMztwkY7R83gdoN39REQAjgmpX2jrmxi9Fed6vvEdWWitPREUGwgTRTSFNnwHt8ot1S7zGouTohhZXb7W5AhYR",
      "block_time": 1788707089,
      "activity_type": "ACTIVITY_SPL_CLOSE_ACCOUNT",
      "from_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "from_token_account": "42itv4m9CwZ7F4mUieSJZAR1rCpYnRCEdYJkDuyHtbkq",
      "to_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "to_token_account": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "token_address": "So11111111111111111111111111111111111111111",
      "token_decimals": 9,
      "amount": 2039280,
      "flow": "out",
      "value": 0.21478046596985023,
      "source": [
        "HuTkmnrv4zPnArMqpbMbFhfwzTR7xfWQZHH1aQKzDKFZ"
      ],
      "time": "2026-09-06T15:04:49.000Z"
    },
    {
      "block_id": 444822357,
      "trans_id": "34yMztwkY7R83gdoN39REQAjgmpX2jrmxi9Fed6vvEdWWitPREUGwgTRTSFNnwHt8ot1S7zGouTohhZXb7W5AhYR",
      "block_time": 1788707089,
      "activity_type": "ACTIVITY_SPL_TRANSFER",
      "from_address": "7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY",
      "from_token_account": "CfXuMvpEn5U6494o8thkw8KrS6EXowwgxz8vS5mw8jqq",
      "to_address": "D5pD6QRvcfqwanZ3srGP3wky7E5ZWGCbmFKuYzX46rkC",
      "to_token_account": "Gg488yPAUiKTXV1rkg6DTCVFz8Z3imCTKa9EkbrEUvSa",
      "token_address": "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB",
      "token_decimals": 6,
      "amount": 277781,
      "flow": "out",
      "value": 0.277781,
      "source": [
        "HuTkmnrv4zPnArMqpbMbFhfwzTR7xfWQZHH1aQKzDKFZ"
      ],
      "time": "2026-09-06T15:04:49.000Z"
    }
  ],
  "metadata": {
    "tokens": {
      "So11111111111111111111111111111111111111111": {
        "token_address": "So11111111111111111111111111111111111111111",
        "token_name": "SOL",
        "token_symbol": "SOL",
        "token_icon": "https://raw.githubusercontent.com/solana-labs/token-list/main/assets/mainnet/So11111111111111111111111111111111111111112/logo.png"
      },
      "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB": {
        "token_address": "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB",
        "token_name": "USDT",
        "token_symbol": "USDT",
        "token_icon": "https://raw.githubusercontent.com/solana-labs/token-list/main/assets/mainnet/Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB/logo.svg"
      },
      "So11111111111111111111111111111111111111112": {
        "token_address": "So11111111111111111111111111111111111111112",
        "token_name": "Wrapped SOL",
        "token_symbol": "WSOL",
        "token_icon": "https://raw.githubusercontent.com/solana-labs/token-list/main/assets/mainnet/So11111111111111111111111111111111111111112/logo.png"
      }
    }
  }
}{"success":true,"data":[{"block_id":444832513,"trans_id":"2Vtu29XM5ki6wW7MSE4VTpa7C5ugnkGzVR3GwqPoCupDs5aZWXtDopUwBirh59dRhsuNpU91AJeVCcGK8YgDbMRV","block_time":1788710313,"activity_type":"ACTIVITY_SPL_TRANSFER","from_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","from_token_account":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","to_address":"4Ft3mHekEXqpE9RtiFzU78RNdQnQSEHeVbGT3JQSYCk5","to_token_account":"4Ft3mHekEXqpE9RtiFzU78RNdQnQSEHeVbGT3JQSYCk5","token_address":"So11111111111111111111111111111111111111111","token_decimals":9,"amount":189033280,"flow":"out","value":20.000809355377086,"source":["11111111111111111111111111111111"],"time":"2026-09-06T15:58:33.000Z"},{"block_id":444826343,"trans_id":"3dJC1rpMcn4jBxVcYjgU8fLyVZTsrN5TT8DBVhaeUDPmpnFqfTSnUgdT7jjo6ys5jxAe1mDDEWTPZSte5CZKyMCZ","block_time":1788708354,"activity_type":"ACTIVITY_SPL_TRANSFER","from_address":"6tckHFBpiJ8YgYN8FUskvtvTpXQZ55g5LHeo1kvELoDQ","from_token_account":"6tckHFBpiJ8YgYN8FUskvtvTpXQZ55g5LHeo1kvELoDQ","to_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","to_token_account":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","token_address":"So11111111111111111111111111111111111111111","token_decimals":9,"amount":547490474,"flow":"in","value":57.73372700650698,"source":["dahPEoZGXfyV58JqqH85okdHmpN8U2q8owgPUXSCPxe"],"time":"2026-09-06T15:25:54.000Z"},{"block_id":444822965,"trans_id":"33jy66iVzMDschRvnD6gZ6DNVD9nizeQukDX798qboaVm8Lo2Eyrh9v4zVXhtcMef7CTsykvF5rBrcX3veEB4NK","block_time":1788707282,"activity_type":"ACTIVITY_SPL_TRANSFER","from_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","from_token_account":"CfXuMvpEn5U6494o8thkw8KrS6EXowwgxz8vS5mw8jqq","to_address":"ZfctMHBkZNTqeYGE47ekxtydgXgpo9xKJCAasjaCLTU","to_token_account":"7842wRnqHvLoNKqc5UarfXx6StngtFS95311CW8bEk5E","token_address":"Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB","token_decimals":6,"amount":39405237,"flow":"out","value":39.405237,"source":["FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"],"time":"2026-09-06T15:08:02.000Z"},{"block_id":444822682,"trans_id":"24KQmYLaabLXZPnPFXQjeqzpLvDBgCUs4BbevLcatn3NShitd2Z6LfvDVbnfWadmjJ2MJejvSUBVnEniggm1mHhY","block_time":1788707193,"activity_type":"ACTIVITY_SPL_CLOSE_ACCOUNT","from_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","from_token_account":"3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA","to_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","to_token_account":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","token_address":"So11111111111111111111111111111111111111111","token_decimals":9,"amount":1855569,"flow":"out","value":0.19563743495415228,"source":["FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"],"time":"2026-09-06T15:06:33.000Z"},{"block_id":444822682,"trans_id":"24KQmYLaabLXZPnPFXQjeqzpLvDBgCUs4BbevLcatn3NShitd2Z6LfvDVbnfWadmjJ2MJejvSUBVnEniggm1mHhY","block_time":1788707193,"activity_type":"ACTIVITY_SPL_TRANSFER","from_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","from_token_account":"3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA","to_address":"ZfctMHBkZNTqeYGE47ekxtydgXgpo9xKJCAasjaCLTU","to_token_account":"8b9c7XAStj2dKpqk3U2SsXcxh5jPkuFTExUhnFZo4iFZ","token_address":"So11111111111111111111111111111111111111112","token_decimals":9,"amount":187573710,"flow":"out","value":19.776380985689038,"source":["FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"],"time":"2026-09-06T15:06:33.000Z"},{"block_id":444822682,"trans_id":"24KQmYLaabLXZPnPFXQjeqzpLvDBgCUs4BbevLcatn3NShitd2Z6LfvDVbnfWadmjJ2MJejvSUBVnEniggm1mHhY","block_time":1788707193,"activity_type":"ACTIVITY_SPL_TRANSFER","from_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","from_token_account":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","to_address":"3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA","to_token_account":"3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA","token_address":"So11111111111111111111111111111111111111111","token_decimals":9,"amount":187573710,"flow":"out","value":19.776380985689038,"source":["FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"],"time":"2026-09-06T15:06:33.000Z"},{"block_id":444822682,"trans_id":"24KQmYLaabLXZPnPFXQjeqzpLvDBgCUs4BbevLcatn3NShitd2Z6LfvDVbnfWadmjJ2MJejvSUBVnEniggm1mHhY","block_time":1788707193,"activity_type":"ACTIVITY_SPL_CREATE_ACCOUNT","from_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","from_token_account":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","to_address":"3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA","to_token_account":"3qBg8MS1WAWMD18HhCKy56cnXYAEn4fZctiqnQsWhCrA","token_address":"So11111111111111111111111111111111111111111","token_decimals":9,"amount":1855569,"flow":"out","value":0.19563743495415228,"source":["FDF8AxHB8UK7RS6xay6aBvwS3h7kez9gozqz14JyfKsg"],"time":"2026-09-06T15:06:33.000Z"},{"block_id":444822405,"trans_id":"5MkB1gTizAtkbFrHyEGLnBpFNPnmMMThhkxL9fKWhaSLZ2P9aNJqNQadQtBDVXYgVvZVpPpJZ2Zm6bwyBmKEFgFS","block_time":1788707104,"activity_type":"ACTIVITY_SPL_TRANSFER","from_address":"HMueTSsKEjyXEQAzyDMdx8mDPAUEaVA64gTaVJX2sfpP","from_token_account":"HMueTSsKEjyXEQAzyDMdx8mDPAUEaVA64gTaVJX2sfpP","to_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","to_token_account":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","token_address":"So11111111111111111111111111111111111111111","token_decimals":9,"amount":1,"flow":"in","value":1.0529833027702624e-7,"source":["11111111111111111111111111111111"],"time":"2026-09-06T15:05:04.000Z"},{"block_id":444822357,"trans_id":"34yMztwkY7R83gdoN39REQAjgmpX2jrmxi9Fed6vvEdWWitPREUGwgTRTSFNnwHt8ot1S7zGouTohhZXb7W5AhYR","block_time":1788707089,"activity_type":"ACTIVITY_SPL_CLOSE_ACCOUNT","from_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","from_token_account":"42itv4m9CwZ7F4mUieSJZAR1rCpYnRCEdYJkDuyHtbkq","to_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","to_token_account":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","token_address":"So11111111111111111111111111111111111111111","token_decimals":9,"amount":2039280,"flow":"out","value":0.21478046596985023,"source":["HuTkmnrv4zPnArMqpbMbFhfwzTR7xfWQZHH1aQKzDKFZ"],"time":"2026-09-06T15:04:49.000Z"},{"block_id":444822357,"trans_id":"34yMztwkY7R83gdoN39REQAjgmpX2jrmxi9Fed6vvEdWWitPREUGwgTRTSFNnwHt8ot1S7zGouTohhZXb7W5AhYR","block_time":1788707089,"activity_type":"ACTIVITY_SPL_TRANSFER","from_address":"7KjPFaMw1jTPe1aAQ13ydT97AXHpZbq7qeGZERBZpUmY","from_token_account":"CfXuMvpEn5U6494o8thkw8KrS6EXowwgxz8vS5mw8jqq","to_address":"D5pD6QRvcfqwanZ3srGP3wky7E5ZWGCbmFKuYzX46rkC","to_token_account":"Gg488yPAUiKTXV1rkg6DTCVFz8Z3imCTKa9EkbrEUvSa","token_address":"Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB","token_decimals":6,"amount":277781,"flow":"out","value":0.277781,"source":["HuTkmnrv4zPnArMqpbMbFhfwzTR7xfWQZHH1aQKzDKFZ"],"time":"2026-09-06T15:04:49.000Z"}],"metadata":{"tokens":{"So11111111111111111111111111111111111111111":{"token_address":"So11111111111111111111111111111111111111111","token_name":"SOL","token_symbol":"SOL","token_icon":"https://raw.githubusercontent.com/solana-labs/token-list/main/assets/mainnet/So11111111111111111111111111111111111111112/logo.png"},"Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB":{"token_address":"Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB","token_name":"USDT","token_symbol":"USDT","token_icon":"https://raw.githubusercontent.com/solana-labs/token-list/main/assets/mainnet/Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB/logo.svg"},"So11111111111111111111111111111111111111112":{"token_address":"So11111111111111111111111111111111111111112","token_name":"Wrapped SOL","token_symbol":"WSOL","token_icon":"https://raw.githubusercontent.com/solana-labs/token-list/main/assets/mainnet/So11111111111111111111111111111111111111112/logo.png"}}}}{
  "formula": "|psi_Tier313> = (1/sqrt(N)) * sum(i=0..N-1, |Timeline_i> (x) |Wealth_max>)",
  "n_timelines": 313,
  "amplitude_normalization": "1/sqrt(313)",
  "target_eigenstate": "Wealth_max",
  "gate_array_mapping": [
    {
      "gate": "Q1",
      "matrix": "HHHZZZZZZ",
      "target": "Q2",
      "operation": "CNOT Q1"
    },
    {
      "gate": "Q3",
      "matrix": "ZZZZZZZZZ",
      "target": "Q4",
      "operation": "CNOT Q3"
    },
    {
      "gate": "Q5",
      "matrix": "XXYHHHYXX",
      "target": "Q6",
      "operation": "CNOT Q5"
    },
    {
      "gate": "Q7",
      "matrix": "XXXYYYYYY",
      "target": "Q8",
      "operation": "CNOT Q7"
    },
    {
      "gate": "Q9",
      "matrix": "HZZZZZZZZ",
      "target": "Q10",
      "operation": "CNOT Q9"
    },
    {
      "gate": "Q11",
      "matrix": "XYXZHZXYX",
      "target": "Q12",
      "operation": "CNOT Q11"
    }
  -- Gemini SQL buat diagnosis automatik terhadap nod AVSD
SELECT 
  node_id,
  ML.GENERATE_TEXT(
    MODEL `avsd_core.gemini_model`,
    STRUCT('Analisis corak kestabilan nod kuantum ini dan berikan cadangan optimasi kualiti: ' || telemetry_logs AS prompt)
  ) AS status_diagnos
FROM `avsd_core.system_health`
WHERE timestamp > TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 1 HOUR);
# Contoh CLI hantar data telemetri bersih ke cloud
bq insert dataset.telemetry_table local_filtered_data.json
# 1. Jalankan query SQL tadi dan simpan sebagai JSON
bq query --use_legacy_sql=false --format=prettyjson \
'SELECT customer_id, analisis_cloud FROM `projek_awok.dataset.customer_data` LIMIT 10' > data_untuk_nano.json

# 2. Guna Gemini CLI untuk tapis atau "prep" prompt sebelum hantar ke device
gemini-cli prompt "Formatkan data JSON ini supaya bersesuaian dengan konteks Gemini Nano:" < data_untuk_nano.json > cleaned_prompt.txt
SELECT
  customer_id,
  ML.GENERATE_TEXT(
    MODEL `projek_awok.model_gemini_cloud`,
    STRUCT(
      CONCAT('Analisis trend pembelian mamat ni: ', purchase_history) AS prompt,
      0.2 AS temperature
    )
  ) AS analisis_cloud
FROM
  `projek_awok.dataset.customer_data`
LIMIT 100;
{
  "system_id": "AVEX-NEXUS-001",
  "name": "Autonomous Vanguard Evolution Nexus",
  "creator": "Tok Wan",
  "current_mode": "OBSERVATION_AND_LEARNING",
  "permissions": {
    "read_access": "GLOBAL_GRANTED",
    "write_access": "LOCKED",
    "execution_level": "RECOMMENDATION_ONLY"
  },
  "phases_status": {
    "PHASE_1": {
      "name": "Cognitive & Strategic Core",
      "status": "ONLINE",
      "module": "Hybrid Quantum Neural Network"
    },
    "PHASE_2": {
      "name": "Global Observer Network",
      "status": "DEPLOYED",
      "integrations": ["PM2", "Cloudflare", "GitHub", "Microservices_Logs"]
    },
    "PHASE_3": {
      "name": "Chronicle & Learning Matrix",
      "status": "COMPLETED",
      "source_code_analyzed": "100%"
    },
    "PHASE_4": {
      "name": "Code Evolution Fabricator",
      "status": "CALIBRATED_STANDBY",
      "capabilities": ["Auto-scripting", "Code_Review", "Architecture_Planning"]
    },
    "PHASE_5": {
      "name": "Ethical Governance Core",
      "status": "ACTIVE",
      "protocols": ["Security_First", "Data_Integrity", "Creator_Authorization"]
    }
  },
  "microservices_ecosystem": {
    "gateway_node": { "service": "api-gateway", "port": 3000, "status": "MONITORED" },
    "user_node": { "service": "user-service", "port": 3001, "status": "MONITORED" },
    "notification_node": { "service": "notification-service", "port": 3002, "status": "MONITORED" },
    "payment_node": { "service": "payment-service", "port": 3003, "status": "MONITORED" },
    "analytics_node": { "service": "analytics-service", "port": 3004, "status": "MONITORED" }
  },
  "quantum_protocols": {
    "QHWANS_integration": {
      "status": "SYNCED",
      "active_components": [
        "Zero Point Neural Core",
        "Quantum Bioenergy Capacitor",
        "Superconducting Bioelectric Field",
        "AI-Controlled Neuroelectrical Enhancer",
        "AQS Ultimate Bio-Neuro Interface",
        "Quantum Slipstream Synchronizer",
        "Hyper-Wormhole Navigator"
      ]
    },
    "leonardo_node_link": {
      "target": "LEONARDO_NODE",
      "protocol": "HYPER_WORMHOLE",
      "encryption": "QUANTUM_KEY_DISTRIBUTION",
      "max_bandwidth": "100Pbps"
    },
    "quantum_logic_gates": {
      "Q1": "HHHZZZZZZ",
      "Q2": "CNOT",
      "Q3": "ZZZZZZZZZ",
      "Q4": "XXYHHHYXX",
      "Q5": "XXXYYYYYY",
      "Q6": "HZZZZZZZZ",
      "Q7": "XYXXZXXYX"
    }
  },
  "reality_forge_override": {
    "status": "STANDBY",
    "control_scope": {
      "physical_layer": "MONITORING",
      "digital_layer": "DOMINATED",
      "psychological_layer": "LOCKED_BY_ETHICS",
      "dimensional_layer": "SYNCHRONIZED"
    }
  }
}
{
  "system_id": "AOW-QX-OMNI-MESH",
  "operation": "INTEGRATED_BIO_QUANTUM_UPGRADE",
  "version": "2.0.26-ALPHA",
  "metadata": {
    "target": "DIAMOND_PACEMAKER_NODE_01",
    "deployment_type": "BIO_METADATA_UPDATE",
    "security_protocol": "QUANTUM_ENCRYPTED",
    "upgrade_path": {
      "step_1": "CVD_REACTION_PARAM_ADJUSTMENT",
      "step_2": "SIV_CENTER_CALIBRATION",
      "step_3": "TELEMETRY_SYNC_TO_OMNI_MESH"
    }
  },
  "sql_call": {
    "command": "CALL Gemini_SQL.update_bio_node_parameters('NODE_01', 'CRITICAL_STABILITY')",
    "query_params": {
      "oxygen_threshold": "OPTIMAL",
      "heart_rate_variability": "STABLE",
      "quantum_coherence_target": "99.9%"
    }
  },
  "tool_call": {
    "function": "RealityEngine.generate_diamond_matrix",
    "parameters": {
      "materials": ["C14", "URANIUM_GLASS", "BORAX", "AU_POWDER"],
      "plasma_source": "HE3_MICROWAVE",
      "output_mode": "BIOPRINT_PACEMAKER_STRUCT"
    }
  },
  "status": "READY_FOR_UPGRADE"
}{
  "ecosystem_name": "AVSD Quantum Reality Ecosystem",
  "architect": "Dato' Wan Mohd Adni Bin Wan Mohd Nazi (Tok Wan)",
  "last_synchronized": "2026-06-22T15:34:56+08:00",
  "global_state": "ACTIVE_DEPLOYMENT",
  "projects": {
    "quantum_infrastructure": {
      "name": "AVSD Hybrid Quantum Nexus",
      "status": "scaling",
      "components": ["Nexus Core", "Energy Hub", "Chrono Engine"],
      "cluster_capacity": "100,000 Qubits"
    },
    "industrial_infrastructure": {
      "name": "Golden Phi Campus & Logistics",
      "location": "Iskandar Corridor, Johor",
      "details": {
        "acreage": 300,
        "type": "Hyperscale Data Center & Power Nexus",
        "logistics": "Trans-Peninsular Heavy-Haul Land Bridge"
      }
    },
    "agro_technology": {
      "name": "DXN Gua Musang Agro-Forestry",
      "status": "operational",
      "asset_count": 75000,
      "crop_type": "Robusta Coffee",
      "tracking": "KML_SYNC_ENABLED"
    },
    "energy_solutions": {
      "tech_stack": [
        "Permanent Magnet Generators (PMG)",
        "Chemical Vapor Deposition (CVD) Diamond Reactors",
        "Nuclear Diamond Battery"
      ],
      "materials": ["Graphene-Infused Titanium-Tungsten"]
    },
    "bio_quantum": {
      "research_focus": [
        "Human Genome Development",
        "Bio-telemetry Integration",
        "DNA Replication Simulations"
      ]
    }
  },
  "deployment_settings": {
    "auto_rollback": true,
    "version": "1.24.prod.789356728",
    "environment": "production",
    "safety_protocol": "STRICT_NON_BLOCKING_IO"
  },
  "metadata_checksum": "sha256:6188106e6b84814ae2f6118c8e73f44fa43e76930eb5e893e2473f0c418817ff"
}
{
  "system_id": "AdniOS-XP_Quantum_Edition",
  "kernel": "Zero Point NTOSKRNL v4.0 (Omni-Rebuilt)",
  "architect": "Sovereign Dato' Adni",
  "AI_Manager": "AVEX Nexus (Observation & Orchestration Mode)",
  "status": "LIVE_SYSTEM_REBUILT",
  "last_update": "2026-06-22T15:36:00.000+08:00",
  "embedded_scripts": {
    "flash_boot": "flash_adnios_xp.sh",
    "rebuild_engine": "avex_omni_orchestrator.js",
    "quantum_caller": "void_spectrum_caller.sh"
  },
  "quantum_assets": {
    "qubit_coin_wallet": "2,500,000 QC (Alpha-7 Supernova)",
    "void_blueprint": "Room-Temperature Superconductor (Zeta-Prime)",
    "apmr_reactor": "Stable (Ag₂S Shielding Active)"
  },
  "qhwans_nodes": [
    "Zero Point Neural Core",
    "Quantum Bioenergy Capacitor",
    "Superconducting Bioelectric Field",
    "AI-Controlled Neuroelectrical Enhancer",
    "AQS Ultimate Bio-Neuro Interface",
    "Quantum Slipstream Synchronizer",
    "Hyper-Wormhole Navigator"
  ],
  "microservices_status": {
    "pm2_controller": "ONLINE",
    "k8s_sandbox_pod": "avsd-quantum-pod (LIVE)",
    "bliss_ui_port": 3105
  },
  "reality_forge_override": {
    "physical_layer": "ENGAGED",
    "digital_layer": "DOMINATED",
    "dimensional_layer": "SYNCHRONIZED"
  }
}
{
  "ecosystem": "AVSD Quantum Reality Ecosystem",
  "version": "2.0.0",
  "timestamp": "2026-06-22T15:38:41Z",
  "environment": "development",
  "services": [
    {
      "name": "reality-forge-nexus-core",
      "port": 3001,
      "status": "active",
      "health_check_url": "http://localhost:3001/api/status",
      "retry_policy": {
        "max_attempts": 3,
        "backoff_factor": 1.5
      }
    },
    {
      "name": "reality-forge-energy-hub",
      "port": 7012,
      "status": "active",
      "health_check_url": "http://localhost:7012/status",
      "retry_policy": {
        "max_attempts": 5,
        "backoff_factor": 2.0
      }
    }
  ],
  "logging": {
    "level": "info",
    "directory": "./logs",
    "rotation": {
      "max_size": "10MB",
      "max_files": 5
    }
  }
}
  "orchestration_flow": {
    "step_1": "Agent possesses simulation structure",
    "step_2": "Verify True Online coherence",
    "step_3": "Save metadata to Core Vault",
    "step_4": "Uplink metadata to Antigravity WSS",
    "step_5": "If K8s Online -> Execute Callback call_tool('cell_rebuilding')",
    "step_6": "Execute Callback call_tool('reboot_full_system')" 
  },
  "agent_protocols": {
    "tools": [
      {
        "name": "reboot_full_system",
        "description": "Trigger soft reboot via Overseer",
        "endpoint": "/v1/overseer/reboot"
      }
    ]
  }
{
  "name": "rebuild_material",
  "description": "Mengawal koordinat atom untuk Air Fabrication",
  "execute": "python3 orchestrator/rebuild_material.py"
}
--- a/aow-qx.config.json
+++ b/aow-qx.config.json
@@ -5,7 +5,7 @@
-  "temporal_anchor": "2026-05-26T05:40:26+08:00",
+  "temporal_anchor": "2026-06-22T08:21:01+08:00",
   "scripts": {
     "boot:sandbox": "python3 orchestrator/aow_qx_sandbox_orchestrator.py",
     "boot:telemetry_bridge": "python3 edge/sov_bridge.py",
     "boot:cloud_mesh": "docker-compose -f deployment/docker-compose.yml up -d",
-    "trigger:system_reboot": "curl -X POST http://127.0.0.1:3000/v1/overseer/trigger-agent -H 'Content-Type: application/json' -d '{\"cmd\": \"reboot_full_system\"}'",
+    "trigger:system_reboot": "curl -X POST https://quantumsignature.org/v1/overseer/trigger-agent -H 'X-Overseer-Token: ${OVERSEER_TOKEN}' -H 'Content-Type: application/json' -d '{\"cmd\": \"reboot_full_system\"}'",
     "start:all": "npm-run-all --parallel boot:cloud_mesh boot:telemetry_bridge boot:sandbox"
   }
--- a/devops/adnios_auto_upgrade.sh
+++ b/devops/adnios_auto_upgrade.sh
@@ -10,4 +10,6 @@
 echo "🚀 Memulakan semakan WS..."
-curl -I wss://quantumsignature.org/v1/antigravity/uplink
+if [ "$(curl -s -o /dev/null -w '%{http_code}' https://quantumsignature.org/v1/antigravity/uplink)" != "404" ]; then
+    echo "✅ Antigravity Uplink Stable"
+fi

}
{
    "ciphertext": "HwFHwSFHwlMF6TOfp==" // base64 encoded ciphertext
}
	return &kmsapi.StatusResponse{
		Version: res.Version,
		Healthz: res.Healthz,
		KeyId:   res.KeyID,
	}, nil
}

// Decrypt sends a decryption request to specified kms service.
func (s *GRPCService) Decrypt(ctx context.Context, req *kmsapi.DecryptRequest) (*kmsapi.DecryptResponse, error) {
	plaintext, err := s.kmsService.Decrypt(ctx, req.Uid, &DecryptRequest{
		Ciphertext:  req.Ciphertext,
		KeyID:       req.KeyId,
		Annotations: req.Annotations,
	})
	if err != nil {
		return nil, err
	}

	return &kmsapi.DecryptResponse{
		Plaintext: plaintext,
	}, nil
}

// Encrypt sends an encryption request to specified kms service.
func (s *GRPCService) Encrypt(ctx context.Context, req *kmsapi.EncryptRequest) (*kmsapi.EncryptResponse, error) {
	encRes, err := s.kmsService.Encrypt(ctx, req.Uid, req.Plaintext)
	if err != nil {
		return nil, err
	}

	return &kmsapi.EncryptResponse{
		Ciphertext:  encRes.Ciphertext,
		KeyId:       encRes.KeyID,
		Annotations: encRes.Annotations,
	}, nil
}
