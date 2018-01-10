/*
 * Requires https://github.com/dgruber/jsv
 */

package main

import (
    "strings"
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
    } else {
        pe_name, _ := jsv.GetParam("pe_name")

        // XXX the "shm" PE is the single-node multicore PE
        //     change this to the equivalent for your site
        // XXX note that this does not properly deal with a range of number of slots;
        //     it just takes the max value of the range
        if strings.EqualFold("shm", pe_name) {
            pe_max, _ := jsv.GetParam("pe_max")
            jsv.SetParam("binding_strategy", "linear_automatic")
            jsv.SetParam("binding_type", "set")
            jsv.SetParam("binding_amount", pe_max)
            jsv.SetParam("binding_exp_n", "0")
            modified_p = true
        }
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

