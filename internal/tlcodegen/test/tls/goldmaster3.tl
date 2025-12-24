---types---
// test field mask var_num and var_bit
aaa  {A:#} {T:Type} {U:Type}  a:# B:# A0:A.0?string A1:A.1?string a0:a.0?string a1:a.1?string B0:B.0?string B1:B.1?U = Aaa A T U;

curl.responseOk#57d7a33a
  field_mask:   #
  http_code:    int
  http_headers: %(Dictionary string)
  http_body:    field_mask.0?string
  = curl.Response;

curl.responseError#af514651
  field_mask:   #
  error_code:   int
  error_string: string
  = curl.Response;

---functions---

@kphp
curl.request#3f5a4651
  field_mask: #
  method:             string
  url:                string
  headers:            field_mask.0?%(Dictionary string)
  body:               field_mask.1?string
  timeout:            field_mask.2?int
  timeout_connection: field_mask.3?int
  max_redirs:         field_mask.4?int
  ssl_version:        field_mask.5?int
  resolve_ipv:        field_mask.6?int
  http_proxy:         field_mask.7?string
  limit:              field_mask.8?int
  route:              field_mask.9?string
  version:            field_mask.10?int
  origin:             field_mask.11?string
  tls_cert:           field_mask.12?string
  tls_cert_pk:        field_mask.13?string
  = curl.Response;
