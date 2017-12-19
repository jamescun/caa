# Authorities.json

Authorities.json contains an array of known certificate authorities, the fingerprints of the certificates they issue against, and the domains they accept for CAA validation.

If you notice an authority is missing, please open a Pull Request and adding following the format defined below.


## Data Format

Authorities.json is an array of objects, where each object represents a legal entity which issues certificates. An authority may trade under a different name or issue certificates on behalf of another entity, these details are rolled into the same object as the primary authority.

Name         | Type          | Required | Description
-------------|---------------|----------|-----------------------------------------------------------
name         | string        | yes      | legal name of certificate authority
aliases      | array[string] | no       | array of alternative names for certificate authority
domains      | array[string] | yes      | array of domain names that are acceptable in a CAA record
certificates | object        | yes      | defines roots/intermediates an authority uses for issuance

### Fingerprint Format

Fingerprints are hexadecimal encoded certificate fingerprints, prefixed by the name of the cryptographic checksum used and a colon separator.

Currently `MD5`, `SHA1` and `SHA256` are expected to be accepted.
