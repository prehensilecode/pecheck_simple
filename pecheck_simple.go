/*
 * Requires https://github.com/dgruber/jsv
 */

package main

import (
    "github.com/dgruber/jsv"
)

func jsv_on_start_function() {
    //jsv_send_env()
}

func job_verification_function() {
    //
    // Set binding on serial jobs (i.e. no PE) to "linear:1
    //
    var modified_p bool = false
    if !jsv.IsParam("pe_name") {
        jsv.SetParam("binding_strategy", "linear_automatic")
        jsv.SetParam("binding_type", "set")
        jsv.SetParam("binding_amount", "1")
        jsv.SetParam("binding_exp_n", "0")
        modified_p = true
    } 

    if modified_p {
        // jsv.JSV_show_params()
        jsv.Correct("Job was modified")
    } else {
        jsv.Correct("Job was not modified")
    }

    return
}

func main() {
    jsv.Run(true, job_verification_function, jsv_on_start_function)
}

