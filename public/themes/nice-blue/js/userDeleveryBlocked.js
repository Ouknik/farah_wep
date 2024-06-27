// init the sweet alert
var _instanceSweetalert = new sweetalert();
// init formData
var formData = {};

function blockedUserDelevery(userdeleveryid) {

    _instanceSweetalert.showBusy("Bitte warten ...");

    formData["userDeleveryId"] = userdeleveryid;
     

    setTimeout(
        function() {
            ajax("/user/delevery/block", "POST", formData, handleSuccessBlockedUserDelevery);
        },500
    );

}


const scenarioSectionsForUpdateElement = document.getElementById("switch-blocked");
function handleSuccessBlockedUserDelevery(responseData) {
    const response = JSON.parse(responseData.response);
    _instanceSweetalert.close();
    // console.log(response)
    if (response.error == 1) {
        setTimeout(function(){
            _instanceSweetalert.init("Der Vorgang wurde erfolgreich abgeschlossen!", response.errorMessage, true, "Ok", true);
            _instanceSweetalert.setType("error");
            _instanceSweetalert.show();
        }, 100);
    } else {

        setTimeout(
            function() {
                _instanceSweetalert.init("Erfolgreich blockiert!", "");
                _instanceSweetalert.setType("success");
                _instanceSweetalert.show();
                setTimeout(function () {
                    _instanceSweetalert.close(1000);
                }, 100);
                // updateInvestorSection_1_4();
                scenarioSectionsForUpdateElement.innerHTML = response.parsedTemplate

            },
        500);
    }
}



