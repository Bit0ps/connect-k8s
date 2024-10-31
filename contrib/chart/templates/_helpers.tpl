{{/* vim: set filetype=mustache: */}}
{{/*
Create blockchain name and version as used by the chart label.
Truncated at 52 chars because StatefulSet label 'blockchain-revision-hash' is limited
to 63 chars and it includes 10 chars of hash and a separating '-'.
*/}}
{{- define "chart.blockchain.fullname" -}}
{{- printf "%s-%s" (include "chart.fullname" .) .Values.blockchain.name | trunc 52 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create the name of the blockchain service account to use
*/}}
{{- define "chart.blockchain.serviceAccountName" -}}
{{- if .Values.blockchain.serviceAccount.create -}}
    {{ default (include "chart.blockchain.fullname" .) .Values.blockchain.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.blockchain.serviceAccount.name }}
{{- end -}}
{{- end -}}

# Renders secrets and envs from template.
{{- define "chart.blockchain.listEnvVars"}}
{{- $fullName := include "chart.blockchain.fullname" . -}}
{{- range $key, $val := .Values.blockchain.extraEnvVars.secret }}
- name: {{ $key }}
  valueFrom:
    secretKeyRef:
      name: {{ $fullName }}-env-secrets
      key: {{ $key }}
{{- end}}
{{- range $key, $val := .Values.blockchain.extraEnvVars.normal }}
- name: {{ $key }}
  value: {{ $val | quote }}
{{- end}}
{{- end }}

{{/*
Create connect name and version as used by the chart label.
*/}}
{{- define "chart.connect.fullname" -}}
{{- printf "%s-%s" (include "chart.fullname" .) .Values.connect.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create the name of the Connect service account to use
*/}}
{{- define "chart.connect.serviceAccountName" -}}
{{- if .Values.connect.serviceAccount.create -}}
    {{ default (include "chart.connect.fullname" .) .Values.connect.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.connect.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Expand the namespace of the release.
Allows overriding it for multi-namespace deployments in combined charts.
*/}}
{{- define "chart.namespace" -}}
{{- default .Release.Namespace .Values.namespaceOverride | trunc 63 | trimSuffix "-" -}}
{{- end }}

# Renders secrets and envs from template.
{{- define "chart.connect.listEnvVars"}}
{{- $fullName := include "chart.connect.fullname" . -}}
{{- range $key, $val := .Values.connect.extraEnvVars.secret }}
- name: {{ $key }}
  valueFrom:
    secretKeyRef:
      name: {{ $fullName }}-env-secrets
      key: {{ $key }}
{{- end}}
{{- range $key, $val := .Values.connect.extraEnvVars.normal }}
- name: {{ $key }}
  value: {{ $val | quote }}
{{- end}}
{{- end }}