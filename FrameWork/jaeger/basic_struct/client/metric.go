package client


// Metrics is a container of all stats emitted by Jaeger tracer.
type Metrics struct {
	//// Number of traces started by this tracer as sampled
	//TracesStartedSampled metrics.Counter `metric:"traces" tags:"state=started,sampled=y" help:"Number of traces started by this tracer as sampled"`
	//
	//// Number of traces started by this tracer as not sampled
	//TracesStartedNotSampled metrics.Counter `metric:"traces" tags:"state=started,sampled=n" help:"Number of traces started by this tracer as not sampled"`
	//
	//// Number of externally started sampled traces this tracer joined
	//TracesJoinedSampled metrics.Counter `metric:"traces" tags:"state=joined,sampled=y" help:"Number of externally started sampled traces this tracer joined"`
	//
	//// Number of externally started not-sampled traces this tracer joined
	//TracesJoinedNotSampled metrics.Counter `metric:"traces" tags:"state=joined,sampled=n" help:"Number of externally started not-sampled traces this tracer joined"`
	//
	//// Number of sampled spans started by this tracer
	//SpansStartedSampled metrics.Counter `metric:"started_spans" tags:"sampled=y" help:"Number of sampled spans started by this tracer"`
	//
	//// Number of unsampled spans started by this tracer
	//SpansStartedNotSampled metrics.Counter `metric:"started_spans" tags:"sampled=n" help:"Number of unsampled spans started by this tracer"`
	//
	//// Number of spans finished by this tracer
	//SpansFinished metrics.Counter `metric:"finished_spans" help:"Number of spans finished by this tracer"`
	//
	//// Number of errors decoding tracing context
	//DecodingErrors metrics.Counter `metric:"span_context_decoding_errors" help:"Number of errors decoding tracing context"`
	//
	//// Number of spans successfully reported
	//ReporterSuccess metrics.Counter `metric:"reporter_spans" tags:"result=ok" help:"Number of spans successfully reported"`
	//
	//// Number of spans not reported due to a Sender failure
	//ReporterFailure metrics.Counter `metric:"reporter_spans" tags:"result=err" help:"Number of spans not reported due to a Sender failure"`
	//
	//// Number of spans dropped due to internal queue overflow
	//ReporterDropped metrics.Counter `metric:"reporter_spans" tags:"result=dropped" help:"Number of spans dropped due to internal queue overflow"`
	//
	//// Current number of spans in the reporter queue
	//ReporterQueueLength metrics.Gauge `metric:"reporter_queue_length" help:"Current number of spans in the reporter queue"`
	//
	//// Number of times the Sampler succeeded to retrieve sampling strategy
	//SamplerRetrieved metrics.Counter `metric:"sampler_queries" tags:"result=ok" help:"Number of times the Sampler succeeded to retrieve sampling strategy"`
	//
	//// Number of times the Sampler failed to retrieve sampling strategy
	//SamplerQueryFailure metrics.Counter `metric:"sampler_queries" tags:"result=err" help:"Number of times the Sampler failed to retrieve sampling strategy"`
	//
	//// Number of times the Sampler succeeded to retrieve and update sampling strategy
	//SamplerUpdated metrics.Counter `metric:"sampler_updates" tags:"result=ok" help:"Number of times the Sampler succeeded to retrieve and update sampling strategy"`
	//
	//// Number of times the Sampler failed to update sampling strategy
	//SamplerUpdateFailure metrics.Counter `metric:"sampler_updates" tags:"result=err" help:"Number of times the Sampler failed to update sampling strategy"`
	//
	//// Number of times baggage was successfully written or updated on spans.
	//BaggageUpdateSuccess metrics.Counter `metric:"baggage_updates" tags:"result=ok" help:"Number of times baggage was successfully written or updated on spans"`
	//
	//// Number of times baggage failed to write or update on spans.
	//BaggageUpdateFailure metrics.Counter `metric:"baggage_updates" tags:"result=err" help:"Number of times baggage failed to write or update on spans"`
	//
	//// Number of times baggage was truncated as per baggage restrictions.
	//BaggageTruncate metrics.Counter `metric:"baggage_truncations" help:"Number of times baggage was truncated as per baggage restrictions"`
	//
	//// Number of times baggage restrictions were successfully updated.
	//BaggageRestrictionsUpdateSuccess metrics.Counter `metric:"baggage_restrictions_updates" tags:"result=ok" help:"Number of times baggage restrictions were successfully updated"`
	//
	//// Number of times baggage restrictions failed to update.
	//BaggageRestrictionsUpdateFailure metrics.Counter `metric:"baggage_restrictions_updates" tags:"result=err" help:"Number of times baggage restrictions failed to update"`
	//
	//// Number of times debug spans were throttled.
	//ThrottledDebugSpans metrics.Counter `metric:"throttled_debug_spans" help:"Number of times debug spans were throttled"`
	//
	//// Number of times throttler successfully updated.
	//ThrottlerUpdateSuccess metrics.Counter `metric:"throttler_updates" tags:"result=ok" help:"Number of times throttler successfully updated"`
	//
	//// Number of times throttler failed to update.
	//ThrottlerUpdateFailure metrics.Counter `metric:"throttler_updates" tags:"result=err" help:"Number of times throttler failed to update"`
}

