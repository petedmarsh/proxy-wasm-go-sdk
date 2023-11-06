## http_body_direct_response

this example demonstrates that OnHttpRequestBody is not called when using a `direct_response``

Run this example, then:

```
$ curl -X POST -d "abc123" localhost:18000
```

You will see logs like:

```
[2023-11-06 22:23:37.315][123297][info][wasm] [source/extensions/common/wasm/context.cc:1171] wasm log: OnHttpRequestHeaders
[2023-11-06 22:23:37.315][123297][info][wasm] [source/extensions/common/wasm/context.cc:1171] wasm log: OnHttpResponseHeaders
[2023-11-06 22:23:37.315][123297][info][wasm] [source/extensions/common/wasm/context.cc:1171] wasm log: OnHttpResponseBody
```

Note that there is no log for OnHttpRequestBody because it was not called.
