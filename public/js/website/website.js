var languagedata
var selectedSlugs = [];
var selectedname = [];
var selectedtemplate

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data

    })


})
$('.hd-crd-btn').click(function () {

    if ($('#hd-crd').is(':visible')) {
        $('#hd-crd').addClass('hidden').removeClass("show");
        document.cookie = `webbanner=false; path=/;`;
    } else {
        $('#hd-crd').addClass("show").removeClass('hidden');
        document.cookie = `webbanner=true; path=/;`;
    }
});


$(document).on('click', '.choosetemplate', function () {

    if ($('.Sitename').val() == "") {
        $('#sitename-error').removeClass('hidden');
        $('#channel-error').addClass('hidden');
        return;
    } else {
        $('#sitename-error').addClass('hidden');
        if (selectedSlugs.length === 0) {
            $('#channel-error').removeClass('hidden');
            return;
        } else {
          
            $('#channel-error').addClass('hidden');

            $.ajax({
                type: 'POST',
                url: '/admin/website/checksitename', // Change to your API endpoint
                data: { "sitename": $('.Sitename').val().trim(),"webid":$('.website_id').val(), csrf: $("input[name='csrf']").val()},
                success: function (response) {
                    console.log("responses",response)
                    if (response) {

                      
                        $('#sitename-error').removeClass('hidden').text("Site name already exists");
                        return;
                    } else {
                        // Proceed if no duplicate
                        $('#sitename-error').addClass('hidden');
                        $('#Id2').removeClass('show');
                        $('#Id3').addClass('show').css('display', 'block');

                        $('.selecttemplate').each(function(){

                             $(this).addClass('hidden')
                             
                          
                           if( $(this).attr('data-type').trim()==$('.channel_name').val().trim()){

                               $(this).removeClass('hidden')

                           }
                        })
                      
                    }
                },
                error: function () {
                    alert("Error checking site name. Please try again.");
                }
            });

            // $('#Id2').removeClass('show');
            // $('#Id3').addClass('show');
            // $('#Id3').css('display', 'block')
        }
    }
});

$(document).on('keyup', '.Sitename', function () {

    if ($(this).val() != "") {

        $('#sitename-error').addClass('hidden')

    }
    $('.site_name').val($(this).val())
})


$(document).on('click', '.channelchkbox', function () {
    var label = $(this);
    var checkbox = $('#' + label.attr('for'));


    setTimeout(function () {
        var slug = label.data('slug');
        var name = label.find('h3').text().trim()
        var isChecked = checkbox.prop('checked');

        if (isChecked) {
            $('#channel-error').addClass('hidden');
            if (selectedSlugs.indexOf(slug) === -1) {
                selectedSlugs.push(slug);
            }
            if (selectedname.indexOf(name) === -1) {
                selectedname.push(name);
            }
        } else {
            selectedSlugs = selectedSlugs.filter(function (item) {
                return item !== slug;
            });
            selectedname = selectedname.filter(function (item) {
                return item !== name;
            });
        }

        $('.channel_name').val(name)
        console.log(selectedSlugs);
    }, 10);
});



$(document).on('click', '.selecttemplate', function () {

    $('#template-error').addClass('hidden')

    $('.selecttemplate.active').removeClass('active');

    $('.selecttemplate').find('.activeimg').addClass('hidden');

    $(this).addClass('active');

    $(this).find('.activeimg').removeClass('hidden')

    selectedtemplate = $(this).data('name');

    console.log(selectedtemplate, "hhhh")

    $('.template_name').val(selectedtemplate)
    $('.template_desc').val($(this).find('.desc').text())
    $('.template_img').val($(this).find('.img').attr('src'))
    $('.template_type').val($(this).data('type'))
    $('#template_id').val('')

    console.log('Selected Template ID:', selectedtemplate);
});


$(document).on('click', '#edit', function () {
    $('.headtag').text('Edit Site')
    $('#Id2').addClass('show')
    $('#Id2').css('display', 'block');
    $('<div class="modal-backdrop fade show"></div>').appendTo(document.body);
    var edit = $(this).closest("tr");
    var sitename = $(this).attr('data-name').trim();
    var channame = edit.find("td:eq(1)").text().trim();
    var templatename = $(this).attr('data-tempname');
 $('.backbtn').removeClass('hidden')
    if (($(this).attr('data-id')==1) &&($(this).attr('data-tenantid')==1)){

        $('.defaultdiv').removeClass('hidden')
        $('.commondiv').addClass('hidden')
    }else{
     $('.defaultdiv').addClass('hidden')
        $('.commondiv').removeClass('hidden')
          

    }

  $('.Sitename').val(sitename);
    $('.site_name').val(sitename)
    $('.template_name').val(templatename)
    $('.channel_name').val(channame)
    $('#webcreatebtn').text("Update")
    $('#createwebsite').attr('action', '/admin/website/updatewebsite')
    $('.website_id').val($(this).attr('data-id'))
    var channels = channame.split(',').map(function (item) {
        return item.trim();
    });

    selectedSlugs.push(channels)
    $('.chandiv input[type=radio]').prop('checked', false).trigger('change');


    channels.forEach(function (channel) {
        $('.chandiv label').each(function () {
            var labelText = $(this).find('h3').text().trim();
            if (labelText === channel) {
                $('#' + $(this).attr('for')).prop('checked', true).trigger('change');
            }
        });
    });
    $('.selecttemplate').each(function () {
        if ($(this).attr('data-name') === templatename) {
            $(this).addClass('active');
            $(this).find('.activeimg').removeClass('hidden')
        } else {
            $(this).removeClass('active');
        }
    })

    selectedname.push(channame)
    console.log("sitenames", sitename, channame, templatename);
});


$(document).on('click', '.cancelbtn', function () {

    $('.headtag').text('Create new site')
    $('.Sitename').val("");
    $('.site_name').val("")
    $('.template_id').val("")
    $('.channel_name').val("")
    $('#channel-error').addClass('hidden')
    $('#sitename-error').addClass('hidden')
    $('#template-error').addClass('hidden')
    $('.selecttemplate').each(function () {
        $(this).removeClass('active');
    })

    $('.activeimg').addClass('hidden')
    $('.chandiv input[type=radio]').prop('checked', false).trigger('change');
    //  $('.chandiv label').prop('checked', false).trigger('change');
    selectedSlugs = [];
    selectedname = [];
    selectedtemplate = ""
    $('#webcreatebtn').text("Create")


    // $('#Id2').removeClass('show');
    $('#Id3').removeClass('show');
    $('#Id3').css('display', '');
    $('#Id2').css('display', 'none');
    $('.modal-backdrop').remove()
    $('#Id2').attr({
        'aria-hidden': 'true',
        'aria-modal': ''
    });
    console.log("selectedSlugs", selectedSlugs)
})


$(document).on('click', '#webcreatebtn', function () {


    if ($('.template_name').val() != "") {
        $('#template-error').addClass('hidden')

        $('#createwebsite').submit()
    } else {

        $('#template-error').removeClass('hidden')
    }

})

$(document).on('click', '.changetemplate', function () {

    $('#Id3').addClass('show')
    $('#Id3').css('display', 'block');
    $('.backbtn').addClass('hidden')
    $('<div class="modal-backdrop fade show"></div>').appendTo(document.body);
    var sitename = $(this).attr('data-name').trim();
    var channame = $(this).attr("data-channame").trim();
    webid = $(this).attr('data-id')
    templateid = $(this).attr('data-tempid')
    templatename = $(this).attr('data-tempname')
    $('.Sitename').val(sitename);
    $('.site_name').val(sitename)
    $('#template_id').val(templateid)
    $('.channel_name').val(channame)
    $('#webcreatebtn').text("Update")
    $('#createwebsite').attr('action', '/admin/website/updatewebsite')
    $('.website_id').val(webid)
    $('.template_name').val(templatename)

    console.log("lll", templateid)
    $('.selecttemplate').each(function () {
        $(this).removeClass('hidden')
        if ($(this).attr('data-name') === templatename) {
            $(this).addClass('active');
            $(this).find('.activeimg').removeClass('hidden')
        } else {
            $(this).removeClass('active');
        }
    })
})

$(document).on('click', '.newsite', function () {

     $('.website_id').val("")
     $('.defaultdiv').addClass('hidden')
        $('.commondiv').removeClass('hidden')
        $('#createwebsite').attr('action','/admin/website/createwebsite')
    $('#Id2').addClass('show')
    $('#Id2').css('display', 'block');
    $('<div class="modal-backdrop fade show"></div>').appendTo(document.body);
     $('.backbtn').removeClass('hidden')

})

$(document).on('click', '.backbtn', function () {


    $('#Id2').addClass('show')
    $('#Id2').css('display', 'block');
    // $('<div class="modal-backdrop fade show"></div>').appendTo(document.body);
    $('#Id3').removeClass('show')
    $('#Id3').css('display', '');
})


$(document).on('click', '.deletebtn', function () {

    var websiteid = $(this).attr("data-id")
    $("#content").text("Are you sure you want to remove this Website")
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')


    if (pageno == null) {
        $('#delid').attr('href', "/admin/website/deletewebsite/" + websiteid );

    } else {
        $('#delid').attr('href', "/admin/website/deletewebsite/" + websiteid + "?page=" + pageno );

    }

   
    $(".deltitle").text("Delete Website ?")
    $('.delname').text($(this).parents('tr').find('td:eq(0) a').text())

})