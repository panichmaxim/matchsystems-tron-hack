module gitlab.com/rubin-dev/api

go 1.18

require (
	github.com/99designs/gqlgen v0.17.7
	github.com/CloudyKit/jet/v6 v6.1.0
	github.com/blocktree/go-owcdrivers v1.2.27
	github.com/cristalhq/jwt/v4 v4.0.0
	github.com/elastic/go-elasticsearch/v8 v8.2.0
	github.com/getsentry/sentry-go v0.13.0
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/gojuno/minimock/v3 v3.0.10
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/neo4j/neo4j-go-driver/v5 v5.0.0-preview
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.8.2
	github.com/rs/zerolog v1.26.1
	github.com/spf13/cast v1.5.0
	github.com/stretchr/testify v1.7.1
	github.com/uptrace/bun v1.1.5
	github.com/uptrace/bun/dbfixture v1.1.5
	github.com/uptrace/bun/dialect/pgdialect v1.1.5
	github.com/uptrace/bun/driver/pgdriver v1.1.5
	github.com/uptrace/bun/extra/bunotel v1.1.5
	github.com/uptrace/uptrace-go v1.7.1
	github.com/vektah/dataloaden v0.3.0
	github.com/vektah/gqlparser/v2 v2.4.3
	gitlab.com/falaleev-golang/config v0.0.0-20220502144359-15b32bd081ee
	gitlab.com/falaleev-golang/mailgate v0.0.0-20220501080148-a73c126178d7
	gitlab.com/falaleev-golang/zlog v0.0.0-20220502144109-373f10bed9fc
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.32.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	golang.org/x/crypto v0.0.0-20220924013350-4ba4fb4dd9e7
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/blocktree/go-owcrypt v1.1.13 // indirect
	github.com/caarlos0/env/v6 v6.9.2 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/drand/kyber v1.1.4 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-chi/chi v4.1.2+incompatible // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailgun/mailgun-go/v3 v3.6.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matryer/moq v0.2.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/phoreproject/bls v0.0.0-20200525203911-a88a5ae26844 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.13 // indirect
	github.com/urfave/cli/v2 v2.8.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.32.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.30.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.30.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.7.0 // indirect
	go.opentelemetry.io/otel/metric v0.30.0 // indirect
	go.opentelemetry.io/otel/sdk v1.7.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.30.0 // indirect
	go.opentelemetry.io/proto/otlp v0.16.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/genproto v0.0.0-20220524164028-0aa58a9287dd // indirect
	google.golang.org/grpc v1.46.2 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
	mellium.im/sasl v0.2.1 // indirect
)
