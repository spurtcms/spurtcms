var languagedata

$(document).ready(async function () {

    var languagecode = $('.language-group>button').attr('data-code')

    await $.getJSON("/locales/"+languagecode+".json", function (data) {
        
        languagedata = data
    })
})