import distro
import platform


class OpSysType(object):
    """Определяет тип ОС с помощью модуля platform"""

    def __getattr__(self, attr):
        """
        Проверяет с каким аттрибутом был получен объект.
        Пример: type = OpSysType => (self == type) => type.lmint, type.rhel
        """
        if attr == "osx":
            return "osx"
        elif attr == "rhel":
            return "redhat"
        elif attr == "centos":
            return "CentOS Linux"
        elif attr == "lmint":
            return "Linux Mint"
        elif attr == "ubnt":
            return "Ubuntu"
        elif attr == "fbsd":
            return "FreeBSD"
        elif attr == "sun":
            return "SunOS"
        elif attr == "unknown_linux":
            return "unknown_linux"
        elif attr == "unknown":
            return "unknown"
        else:
            raise(AttributeError, attr)

    def linuxType(self):
        """Определяет разновидность Linux с помощью различных методов"""
        if distro.linux_distribution()[0] == self.rhel:
            return self.rhel
        elif distro.linux_distribution()[0] == self.lmint:
            return self.lmint
        elif distro.linux_distribution()[0] == self.ubnt:
            return self.ubnt
        elif distro.linux_distribution()[0] == self.centos:
            return self.centos
        else:
            return self.unknown_linux

    def queryOS(self):
        if platform.system() == "Darwin":
            return self.osx
        elif platform.system() == "Linux":
            return self.linuxType()
        elif platform.system() == self.sun:
            return self.sun
        elif platform.system() == self.fbsd:
            return self.fbsd

    def fingerprint():
        ostype = OpSysType()
        print(ostype.queryOS())


if __name__ == "__main__":
    OpSysType.fingerprint()
