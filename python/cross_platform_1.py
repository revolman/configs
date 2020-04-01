import platform
import distro

profile = [
    platform.architecture(),
    distro.linux_distribution(),
    platform.libc_ver(),
    platform.mac_ver(),
    platform.machine(),
    platform.node(),
    platform.platform(),
    platform.processor(),
    platform.python_build(),
    platform.python_compiler(),
    platform.python_version(),
    platform.system(),
    platform.uname(),
    platform.version(),
]

for item in profile:
    print(item)
