<a id="jit-configuration"></a>

## Configuration


 The configuration variable [jit](../server-configuration/query-planning.md#guc-jit) determines whether JIT compilation is enabled or disabled. If it is enabled, the configuration variables [jit_above_cost](../server-configuration/query-planning.md#guc-jit-above-cost), [jit_inline_above_cost](../server-configuration/query-planning.md#guc-jit-inline-above-cost), and [jit_optimize_above_cost](../server-configuration/query-planning.md#guc-jit-optimize-above-cost) determine whether JIT compilation is performed for a query, and how much effort is spent doing so.


 [jit_provider](../server-configuration/client-connection-defaults.md#guc-jit-provider) determines which JIT implementation is used. It is rarely required to be changed. See [Pluggable JIT Providers](extensibility.md#jit-pluggable).


 For development and debugging purposes a few additional configuration parameters exist, as described in [Developer Options](../server-configuration/developer-options.md#runtime-config-developer).
