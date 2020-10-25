# world-time (wt)

wt is an example of how normal users can use custom written SUID programs (owned by root) to gain root privileges on Linux systems. When executed without arguments, wt displays some times from around the world then exits. When executed with arguments, it runs those arguments. SUID programs owned by root are dangerous. They can be used by normal users to compromise or maintain administrative control of previously compromised Linux systems.

## Installation

Normal users may install wt themselves. /home/user/bin is normally a good spot and is the default in the Makefile, but anywhere will do.

```bash
make
make install
```

## Finding SUID Files

On most Linux systems, the find command can be used to locate SUID files. Any SUID file that is not expected should be scrutinized and potentially chmoded to something safer (755) or removed from the system.

```bash
find /usr -perm /4000 2> /dev/null
```

## Notes

  * wt is intended for educational purposes only.
  * Do not use wt on systems without the full knowledge and consent of the system's owner.
  * 
