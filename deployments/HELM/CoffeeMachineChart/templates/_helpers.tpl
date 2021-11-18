{{/*
Expand the name of the chart.
*/}}
{{- define "CoffeeMachine.name" -}}
{{- default .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "CoffeeMachine.fullname" -}}
{{- $name := default .Chart.Name }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "CoffeeMachine.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "CoffeeMachine.labels" -}}
helm.sh/chart: {{ include "CoffeeMachine.chart" . }}
{{ include "CoffeeMachine.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "CoffeeMachine.selectorLabels" -}}
app.kubernetes.io/name: {{ include "CoffeeMachine.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
workload.user.cattle.io/workloadselector: {{ .Release.Name }}-{{ include "CoffeeMachine.name" . }}
{{- end }}