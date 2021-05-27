# Shellcode-Loader-ByteMap

The tool has two functionalities:
1. Encodes the shellcode payload using a ByteMap of the file (key file). The key files could be the compiled Golang project itself.
2. Loads the encoded shellcode into memory and runs it on the victim computer.

Use it to bypass AV's. Support only Windows as of now.
