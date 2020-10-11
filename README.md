# installpgadmin3
A clear way of installing postgresql's pgamin3 in Parrot Security OS.

Import the repository key from [https://www.postgresql.org/media/keys/ACCC4CF8.asc](https://www.postgresql.org/media/keys/ACCC4CF8.asc):

```bash
sudo apt-get install curl ca-certificates gnupg
curl https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
```

Create /etc/apt/sources.list.d/pgdg.list. The distributions are called codename-pgdg. (You may determine the codename of your distribution by running lsb_release -c). In debian, use buster as the actual distribution you are using.For a shorthand version of the above, presuming you are using a supported release:
```bash 
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
```
if it results to an error,change to buster by trying the following :
```bash 
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt buster-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
```

Finally, update the package lists, and start installing packages:

```bash
sudo apt-get update
sudo apt-get install pgadmin3
```

Then start it:
> $ pgadmin3

Have a nice day!
