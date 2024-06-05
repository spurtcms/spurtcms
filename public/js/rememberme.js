function delete_cookie(name) {
    document.cookie = name + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

let RememberFlg = sessionStorage.getItem("rememberme");

let localremFlg = localStorage.getItem("rememberme");

console.log(RememberFlg,localremFlg,"---");

if(RememberFlg == null && localremFlg == null){

    delete_cookie("myspurtcms")

    delete_cookie("myspurtcms2")

    delete_cookie("REMEMBER_ME")

    window.location.reload()

}
