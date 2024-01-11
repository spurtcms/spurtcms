
<p align="center">
  <a href="https://www.spurtcms.com/#gh-light-mode-only">
    <img src="https://www.spurtcms.com/img/spurtcms-logo.svg" width="318px" alt="Spurtcms logo" />
  </a>
   
</p>
<h3 align="center">Open Source Multi Vendor Marketplace for CMS Solution - Self hosted </h3>
<p align="center"> Build with Golang + PostgreSQL</p>
<p align="center"><a href="https://www.spurtcms.com/price-details"> Support PRO</a> · <a href="https://www.spurtcms.com/price-details"> Frontend </a> ·  <a href="https://www.spurtcms.com/price-details"> API Suite </a></p>
<br />
<p align="center">
  <a href="https://github.com/spurtcms/multivendor-marketplace/releases">
    <img src="https://img.shields.io/github/last-commit/spurtcms/deployment" alt="GitHub last commit" />
  </a>
  <a href="https://github.com/spurtcms/multivendor-marketplace/issues">
    <img src="https://img.shields.io/github/issues/spurtcms/deployment" alt="GitHub issues" />
  </a>

  <a href="https://github.com/spurtcms/multivendor-marketplace/releases">
    <img src="https://img.shields.io/github/repo-size/spurtcms/deployment?color=orange" alt="GitHub repo size" />
  </a>
</p>
<br />

> [!IMPORTANT]
> 🎉 <strong>Spurtcms 1.0 is now available!</strong> Read more in the <a target="_blank" href="https://www.spurtcms.com/spurtcms-change-log" rel="dofollow"><strong>announcement post</strong></a>.
<br />

## ❯  🚀 Easy to Deploy Spurtcms Admin Panel on your server

This is the official repository of Spurtcms. Using these Build , you can easily deploy spurtcms Multi-Vendor Marketplace in your local server.

### Step 1:
Navigate to the cloned repository directory “admin-panel” in the terminal and locate the "api" folder


### Step 2:

Navigate to multivendor-marketplace/api folder and Install node_modules  by executing the following command
```
$ npm install
```

It will take few mins for the npm installation to get finished and once done you will see the completion notification messages in terminal.

### Step 3:
Retrieve the "spurtcms_marketplace.sql" file from the "/api" folder and import it into your MySQL server.
### Step 4:
Configure the database settings in the ".env" file located in the "/api" folder, with the name and credentials for the application to connect to your database (imported from spurtcms_marketplace.sql)
 
Database Configuration
we are using MySQL database, we need to configure database credentials in the .env file 

```
#
# MySQL DATABASE
#
TYPEORM_CONNECTION=mysql
TYPEORM_HOST=localhost
TYPEORM_PORT=3306
TYPEORM_USERNAME= "testuser"             #--Your MySql Username
TYPEORM_PASSWORD= "spurt123&"		#--Your MySql Password 
TYPEORM_DATABASE= "spurt_commerce"	#--Your Database Name
TYPEORM_SYNCHRONIZE=false
TYPEORM_LOGGING=["query", "error"]
TYPEORM_LOGGER=advanced-console
```

### Step 5:
In terminal, Navigate to multivendor-marketplace/api folder and Start API execution using the following command:
```
$ node dist/src/app.js
```

## ❯  🚀 Deploy Frontend Admin , Vendor and Store (Angular)


### Step 1:

Navigate to "/var/www/html" (assuming Apache installation has created this directory) from your home directory in your local or server

### Step 2:

*  Copy the "vendor" and "admin" folders as-is directly from "multivendor-marketplace/frontend/" to "/var/www/html/".

*  Copy all folders & files of “store” folder from multivendor-marketplace/frontend/ folder and paste it directly into /var/www/html/

Completion of above steps should successfully setup frontend builds of all 3 panels of spurtcms Marketplace solution such as Store Panel, Vendor Panel and Admin Panel.

* marketplace website is ready to use from  http://{your-domian or IP} (or) http://localhost/
* Vendor Panel can be accessed by http://{your-domian or IP}/vendor/#/auth/login 
* Admin panel be accessed by http://{your-domian or IP}:{your-port}/admin/#/auth/login

Above steps concludes successful installation and setup of spurtcms Marketplace solution build in your local (or) server.


## 🤔 Support , Document and Help

spurtcms 4.8.2 is published to npm under the `@spurtcms/*` namespace.

You can find our extended documentation on our [www.spurtcms.dev](https://www.spurtcms.dev), but some quick links that might be helpful:

- Read [Technology](https://www.spurtcms.com/opensource-ecommerce-multivendor-nodejs-react-angular) to learn about our vision and what's in the box.

- Our [Discard](https://discord.com/invite/hyW4MXXn8n) Questions, Live Discussions [spurtcms Support](https://accounts.spurtcms.com/#/auth/login-client).
- An [API Reference](https://www.spurtcms.dev/v/spurtapi/) contains the details on spurtcms foundational building blocks.
- Some [Video](https://www.youtube.com/@spurtcms/videos) Video Tutorials 
- Every [Release](https://github.com/spurtcms/multivendor-marketplace/releases) is documented on the Github Releases page.

🐞 If you spot a bug, please [submit a detailed issue](https://github.com/spurtcms/multivendor-marketplace/issues/new), and wait for assistance.

🤔 If you have a question or feature request, please [start a new discussion](https://github.com/orgs/spurtcms/discussions/new/choose). 


## ❯ Maintainers
spurtcms is developed and maintain by [Piccosoft Software Labs India (P) Limited,](https://www.piccosoft.com).


## ❯ License

spurtcms is released under the [BSD-3-Clause License.](https://github.com/spurtcms/spurtcms/blob/master/LICENSE).



