# -*- coding: utf-8 -*-


import os


from hashlib import md5, sha1, sha256


from zlib import crc32


import time


from math import ceil










class Hash:


    def __init__(self, strFilePath):


        self.strFilePath = strFilePath  # 文件绝对路径


        self.hashMD5 = ""  # MD5值


        self.hashSHA1 = ""  # SHA1值


        self.hashSHA256 = ""  # SHA256值


        self.hashCRC32 = ""  # CRC32值






    def getMd5(self):  # 计算md5


        mdfive = md5()


        with open(self.strFilePath, 'rb') as f:


            mdfive.update(f.read())


        self.hashMD5 = mdfive.hexdigest().upper()


        return self.hashMD5






    def getSha1(self):  # 计算sha1


        sha1Obj = sha1()


        with open(self.strFilePath, 'rb') as f:


            sha1Obj.update(f.read())


        self.hashSHA1 = sha1Obj.hexdigest().upper()


        return self.hashSHA1






    def getSha256(self):  # 计算sha256


        sha256Obj = sha256()  # Get the hash algorithm.


        with open(self.strFilePath, 'rb') as f:


            sha256Obj.update(f.read())  # Hash the data.


        self.hashSHA256 = sha256Obj.hexdigest().upper()  # Get he hash value.


        return self.hashSHA256






    def getCrc32(self):  # 计算crc32


        with open(self.strFilePath, 'rb') as f:


            self.hashCRC32 = hex(crc32(f.read()))[2:].upper()


        return self.hashCRC32






    def getHash(self):  # 计算Hash


        mdfive = md5()


        sha1Obj = sha1()


        sha256Obj = sha256()






        with open(self.strFilePath, 'rb') as f:


            fileContent = f.read()


            mdfive.update(fileContent)


            self.hashMD5 = mdfive.hexdigest().upper()






            sha1Obj.update(fileContent)


            self.hashSHA1 = sha1Obj.hexdigest().upper()






            sha256Obj.update(fileContent)


            self.hashSHA256 = sha256Obj.hexdigest().upper()






            self.hashCRC32 = hex(crc32(fileContent))[2:].upper()






    def printHash1(self):


        print('{:8} {}'.format('MD5:', self.getMd5()))


        print('{:8} {}'.format('SHA1:', self.getSha1()))


        print('{:8} {}'.format('SHA256:', self.getSha256()))


        print('{:8} {}'.format('CRC32:', self.getCrc32()))






    def printHash2(self):


        self.getHash()


        print('{:8} {}'.format('MD5:', self.hashMD5))


        print('{:8} {}'.format('SHA1:', self.hashSHA1))


        print('{:8} {}'.format('SHA256:', self.hashSHA256))


        print('{:8} {}'.format('CRC32:', self.hashCRC32))






    # 计算文件大小（以格式KB、MB、GB、TB显示）


    def size_format(self):


        size = os.path.getsize(self.strFilePath)


        # print('文件大小:{}字节'.format(size))






        if size < 1024:


            return '%i' % size + 'size'


        elif 1024 <= size < (1024 * 1024):


            return '%.2f' % float(size / 1024) + 'KB'


        elif (1024 * 1024) <= size < (1024 * 1024 * 1024):


            return '%.2f' % float(size / (1024 * 1024)) + 'MB'


        elif (1024 * 1024 * 1024) <= size < (1024 * 1024 * 1024 * 1024):


            return '%.2f' % float(size / (1024 * 1024 * 1024)) + 'GB'


        elif (1024 * 1024 * 1024 * 1024) <= size:


            return '%.2f' % float(size / (1024 * 1024 * 1024 * 1024)) + 'TB'


        else:


            pass






    def secondsToStr(self, seconds):


        x = time.localtime(seconds)  # 时间元组


        return time.strftime("%Y-%m-%d %X", x)  # 时间元组转为字符串






    '''


    st_mode=33206 #权限模式


    st_ino=844424930150465 #inode number


    st_dev=3795105997 #device


    st_nlink=1 #number of hard links


    st_uid=0  #所有用户的user id


    st_gid=0 #所有用户的group id


    st_size=64985 #文件的大小，以字节为单位


    st_atime=1549040523  #文件最后访问时间


    st_mtime=1549040524  #文件最后修改时间


    st_ctime=1549036862  #文件创建时间


    '''


    def printFileAttributes(self):


        print('文件路径:{}'.format(self.strFilePath))






        # 获取文件属性信息


        fileInfo = os.stat(self.strFilePath)


        print('文件大小:{}'.format(self.size_format()))


        print('文件大小:{}字节'.format(ceil(fileInfo.st_size)))


        print('文件创建时间:{}'.format(self.secondsToStr(fileInfo.st_ctime)))


        # print('文件访问时间:{}'.format(self.secondsToStr(fileInfo.st_atime)))


        print('文件修改时间:{}'.format(self.secondsToStr(fileInfo.st_mtime)))










"""-----------------------------------------------


主函数


-----------------------------------------------"""


if __name__ == '__main__':


    strFileAbsolutePath = os.path.join("../frontend/doutok/public/logo.png")
    # strFileAbsolutePath = os.path.join("./total.pdf")






    objHash = Hash(strFileAbsolutePath)


    # objHash.printHash1()


    objHash.printHash2()


    objHash.printFileAttributes()


