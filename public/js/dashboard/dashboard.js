var languagedata

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })
})
$(document).ready(function () {

    $('.categorynametd').each(function () {
  
      var length = $(this).children('.categorysname').length - 1
  
      $(this).children('.categorysname').each(function (index) {
  
        if (length == index) {
  
          $(this).next().remove();
        }
  
      })
  
    })
})