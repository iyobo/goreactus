
module.exports = {
    login(username, pass, cb) {
        cb = arguments[arguments.length - 1]
        if (localStorage.token) {
            if (cb) cb(true)
            this.onChange(true)
            return
        }
        authAjaxRequest(username, pass, (res) => {
            if (res.authenticated) {
                localStorage.token = res.token
                localStorage.username = res.username
                localStorage.userid = res.userid
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

    getUsername(){
        return localStorage.username
    },

    logout(cb) {
        delete localStorage.token
        delete localStorage.user
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
                username: data.username,
                userid: data.userid
            })
        }
        else{
            cb({ authenticated: false })
        }
    });

}