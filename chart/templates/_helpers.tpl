{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "base.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "base.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" $name .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "base.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*  Manage the labels for each entity  */}}
{{- define "base.labels" -}}
helm.sh/chart: {{ include "base.chart" . }}
{{ include "base.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- range $key, $val := .Values.additionalLabels }}
{{ $key }}: {{ $val | quote }}
{{- end -}}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "base.selectorLabels" -}}
app.kubernetes.io/name: {{ include "base.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "base.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "base.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Generate podAntiAffinity spec
*/}}
{{- define "base.antiAffinityPreset" -}}
{{- $type := .Values.affinity.podAntiAffinityPreset.type -}}
{{- $topologyKey := printf "kubernetes.io/%s" .Values.affinity.podAntiAffinityPreset.topologyKey -}}
podAntiAffinity:
  {{- if eq $type "hard" }}
  requiredDuringSchedulingIgnoredDuringExecution:
    - topologyKey: {{ $topologyKey | quote }}
      labelSelector:
        matchLabels:
          {{- include "base.selectorLabels" . | nindent 10 }}
  {{- else if eq $type "soft" }}
  preferredDuringSchedulingIgnoredDuringExecution:
    - weight: 100
      podAffinityTerm:
        topologyKey: {{ $topologyKey | quote }}
        labelSelector:
          matchLabels:
            {{- include "base.selectorLabels" . | nindent 12 }}
  {{- end }}
{{- end -}}

{{/*
base.rawMetadata will create a resource template that can be
merged with each item in `.Values.raw.resources` and `.Values.raw.templates`.
*/}}
{{- define "base.rawMetadata" -}}
metadata:
  labels:
    {{- include "base.labels" . | nindent 4 }}
{{- end }}
