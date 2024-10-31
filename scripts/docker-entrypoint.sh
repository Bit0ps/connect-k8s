#!/bin/sh

set -e

MARKETMAP_PROVIDER=${MARKETMAP_PROVIDER:-dydx_api}
MARKETMAP_ENDPOINT=${MARKETMAP_ENDPOINT:-dydx-api.polkachu.com:443}
RUN_PPROF=${RUN_PPROF:-true}
PPROF_PORT=${PPROF_PORT:-6060}
METRICS_ENABLED=${METRICS_ENABLED:-false}
LOG_DISABLE_FILE_ROTATION=${LOG_DISABLE_FILE_ROTATION:-false}
LOG_FILE_DISABLE_COMPRESSION=${LOG_FILE_DISABLE_COMPRESSION:-false}
ORACLE_CONFIG=${ORACLE_CONFIG:-""}
LOG_STDOUT_LEVEL=${LOG_STDOUT_LEVEL:-info}
LOG_FILE_LEVEL=${LOG_FILE_LEVEL:-info}
LOG_FILE=${LOG_FILE:-sidecar.log}
LOG_MAX_SIZE=${LOG_MAX_SIZE:-100}
LOG_MAX_BACKUPS=${LOG_MAX_BACKUPS:-1}
LOG_MAX_AGE=${LOG_MAX_AGE:-3}
METRICS_PROMETHEUS_ADDRESS=${METRICS_PROMETHEUS_ADDRESS:-0.0.0.0:8002}
MODE=${MODE:-exec}
HOST=${HOST:-0.0.0.0}
PORT=${PORT:-8080}
UPDATE_INTERVAL=${UPDATE_INTERVAL:-250000000}
MAX_PRICE_AGE=${MAX_PRICE_AGE:-2m0s}

# Set arguments for connect command
CONNECT_ARGS="--marketmap-provider ${MARKETMAP_PROVIDER} \
              --market-map-endpoint ${MARKETMAP_ENDPOINT} \
              --log-std-out-level ${LOG_STDOUT_LEVEL} \
              --log-file-level ${LOG_FILE_LEVEL} \
              --log-file ${LOG_FILE} \
              --log-max-size ${LOG_MAX_SIZE} \
              --log-max-backups ${LOG_MAX_BACKUPS} \
              --log-max-age ${LOG_MAX_AGE} \
              --mode ${MODE} \
              --host ${HOST} \
              --port ${PORT} \
              --update-interval ${UPDATE_INTERVAL} \
              --max-price-age ${MAX_PRICE_AGE}"


# Conditionally add --run-pprof and --pprof-port if RUN_PPROF is set to true
if [ "${RUN_PPROF}" = "true" ]; then
    CONNECT_ARGS="${CONNECT_ARGS} --run-pprof --pprof-port ${PPROF_PORT}"
fi

# Conditionally add --metrics-enabled and --metrics-prometheus-address if METRICS_ENABLED is set to true
if [ "${METRICS_ENABLED}" = "true" ]; then
    CONNECT_ARGS="${CONNECT_ARGS} --metrics-enabled --metrics-prometheus-address ${METRICS_PROMETHEUS_ADDRESS}"
fi

# Conditionally add --log-disable-file-rotation if LOG_DISABLE_FILE_ROTATION is set to true
if [ "${LOG_DISABLE_FILE_ROTATION}" = "true" ]; then
    CONNECT_ARGS="${CONNECT_ARGS} --log-disable-file-rotation"
fi

# Conditionally add --log-disable-file-rotation if LOG_FILE_DISABLE_COMPRESSION is set to true
if [ "${LOG_FILE_DISABLE_COMPRESSION}" = "true" ]; then
    CONNECT_ARGS="${CONNECT_ARGS} --log-file-disable-compression"
fi

# Conditionally add --oracle-config if ORACLE_CONFIG is not empty
if [ -n "${ORACLE_CONFIG}" ]; then
    CONNECT_ARGS="${CONNECT_ARGS} --oracle-config ${ORACLE_CONFIG}"
fi

# Run connect command with arguments
connect ${CONNECT_ARGS}

# Execute any additional commands passed to the script
exec "$@"