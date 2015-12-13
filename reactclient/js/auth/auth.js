
module.exports = {
    login: function(username, pass, cb) {
    cb = arguments[arguments.length - 1]
    if (localStorage.token) {
        if (cb) cb(true)
        this.onChange(true)
        return
    }
    authAjaxRequest(username, pass, (res) => {
        if (res.authenticated) {
        localStorage.token = res.token
        if (cb) cb(true)
            this.onChange(true)
        } else {
            if (cb) cb(false)
            this.onChange(false)
        }
    })
},

getToken() {
    return localStorage.token
},

logout(cb) {
    delete localStorage.token
    if (cb) cb()
    this.onChange(false)
},

loggedIn() {
    return !!localStorage.token
},

onChange() {}
}

function authAjaxRequest(username, password, cb) {

    $.post(SERVER_URL+"/v1/user/login",{username: username, password: password}, function(response){
        var data = JSON.parse(response);
        if(data && data.token){
            cb({
                authenticated: true,
                token: data.token,
                name: data.username
            })
        }
        else{
            cb({ authenticated: false })
        }
    });

}