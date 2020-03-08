# gowhois

## Usage:
```
cat domain.txt | gowhois
echo 'google.com' | gowhois
```

## Output:
```
{
  "domain": {
    "id": "2138514_DOMAIN_COM-VRSN",
    "domain": "GOOGLE.COM",
    "name": "google",
    "extension": "com",
    "status": "clientdeleteprohibited,clienttransferprohibited,clientupdateprohibited,serverdeleteprohibited,servertransferprohibited,serverupdateprohibited",
    "dnssec": "unsigned",
    "whois_server": "whois.markmonitor.com",
    "name_servers": "ns1.google.com,ns2.google.com,ns3.google.com,ns4.google.com",
    "created_date": "1997-09-15T04:00:00Z",
    "updated_date": "2019-09-09T15:39:04Z",
    "expiration_date": "2028-09-14T04:00:00Z"
  },
  "registrar": {
    "id": "292",
    "name": "MarkMonitor Inc.",
    "organization": "",
    "street": "",
    "city": "",
    "province": "",
    "postal_code": "",
    "country": "",
    "phone": "+1.2083895740",
    "phone_ext": "",
    "fax": "",
    "fax_ext": "",
    "email": "abusecomplaints@markmonitor.com",
    "referral_url": "http://www.markmonitor.com"
  },
  "registrant": {
    "id": "",
    "name": "",
    "organization": "",
    "street": "",
    "city": "",
    "province": "",
    "postal_code": "",
    "country": "",
    "phone": "",
    "phone_ext": "",
    "fax": "",
    "fax_ext": "",
    "email": "",
    "referral_url": ""
  },
  "administrative": {
    "id": "",
    "name": "",
    "organization": "",
    "street": "",
    "city": "",
    "province": "",
    "postal_code": "",
    "country": "",
    "phone": "",
    "phone_ext": "",
    "fax": "",
    "fax_ext": "",
    "email": "",
    "referral_url": ""
  },
  "technical": {
    "id": "",
    "name": "",
    "organization": "",
    "street": "",
    "city": "",
    "province": "",
    "postal_code": "",
    "country": "",
    "phone": "",
    "phone_ext": "",
    "fax": "",
    "fax_ext": "",
    "email": "",
    "referral_url": ""
  },
  "billing": {
    "id": "",
    "name": "",
    "organization": "",
    "street": "",
    "city": "",
    "province": "",
    "postal_code": "",
    "country": "",
    "phone": "",
    "phone_ext": "",
    "fax": "",
    "fax_ext": "",
    "email": "",
    "referral_url": ""
  }
}

```