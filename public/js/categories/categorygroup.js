var languagedata
var selectedcheckboxarr = []
/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
        })

            $('.Content').addClass('checked');

})

$(document).ready(function () {
    var $textarea = $('#category_desc');
    var $errorMessage = $('#error-messagedesc');
    var maxLength = 250;

    $textarea.on('input', function () {
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorMessage.text(languagedata.Permission.descriptionchat);
        } else {
            // Clear error message if under the limit
            $errorMessage.text('');
        }
    });

});

$(document).ready(function () {

    var $inputfield = $('#category_name');
    var $errorMessagename = $('#error-messagename');
    var maxLength = 50;

    $inputfield.on('input', function () {
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorMessagename.text("Maximum 50 character allowed");
        } else {
            // Clear error message if under the limit
            $errorMessagename.text('');
        }
    });
})

$(document).ready(function () {
    $('#category_desc').on('input', function () {

        let lines = $(this).val().split('\n').length;

        if (lines > 5) {

            let value = $(this).val();

            let linesArray = value.split('\n').slice(0, 5);

            $(this).val(linesArray.join('\n'));
        }
    });
});



$("#add-btn ,#clickadd").click(function () {
    $("#category-title").text(languagedata.Categoryy.addnewcategorygrp)
    $(".input-group").removeClass("input-group-error")

    $("#category_name-error").hide()
    $("#category_desc-error").hide()
    $("input[name=category_name]").val("");
    $("textarea[name=category_desc]").text("");
    $("#update-btn").hide()
    $("#save").show()
    $("#cancel").show()

})


var edit;
$("body").on("click", "#edit", function () {

    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')

    $(".input-group").removeClass("input-group-error")


    $("#category-title").text(languagedata.Categoryy.updatecatgrp)
    var data = $(this).attr("data-id");
    edit = $(this).closest("tr");
    $("#category_form").attr("name", "editcategory").attr("action", "/categories/updatecategory")
    var name = edit.find("td:eq(1)").text();
    var desc = edit.find("td:eq(2)").text();
    $("input[name=category_name]").val(name.trim());
    $("textarea[name=category_desc]").text(desc.trim());
    $("input[name=category_id]").val(data);
    $("#save").hide()
    $("#update-btn").show()
    $("#cancel").show()

    $("#category_name-error").hide()
    $("#category_desc-error").hide()
    $('#catpageno').val(pageno)
});

$(document).on('click', '.cancel-btn', function () {
    $("#add-btn").show()
    $("#update-btn").hide()
    $("#cancel-btn").hide()
    $("#category_form").attr("name", "createcategory").attr("action", "/categories/newcategory")
    $("input[name=category_id]").val("")
    $("input[name=category_name]").val("")
    $("input[name=category_desc]").val("")
    $("#category_name-error").hide()
    $("#category_desc-error").hide()
    $("#error-messagename").html("")
    $("#error-messagedesc").html("")


})

/*search */
$(document).on("click", "#filterformsubmit", function () {
    var key = $(this).siblings().children(".search").val();
    if (key == "") {
        window.location.href = "/categories/"
    } else {
        $('.filterform').submit();
    }
})

$(document).on('click', '#delete-btn', function () {

    var categoryId = $(this).attr("data-id")
    $("#content").text(languagedata.Categoryy.delcatgrp)
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')

    if (pageno == null) {
        $('#delid').attr('href', "/categories/deletecategory/" + categoryId);

    } else {
        $('#delid').attr('href', "/categories/deletecategory/" + categoryId + "?page=" + pageno);

    }
    $(".deltitle").text(languagedata.Categoryy.delgrptitle)
    $('.delname').text($(this).parents('tr').find('td:eq(1)').text())

})

$("#popup").on("click", function () {

    $("#save").show();
    $("#update-btn").hide();
    $("#cancel-btn").hide();
    $("input[name=category_name]").val("");
    // $("input[name=category_desc]").val("");
    $("textarea[name=category_desc]").text("");
    $("label.error").hide();

    $("#title").text(languagedata.Categoryy.addnewcategorygrp)
})

// Ajax call for sub category list

// Sub Category List In Modal View

$(document).on("click", "#categorylist", function () {
    var id = $(this).attr("data-id")
    var text = $(this).text()
    var url = $("#searchforminclist").attr("action", "/categories/filtercategory/" + id)
    $.ajax({
        url: "/categories/addcategory/" + id,
        type: "POST",
        datatype: "json",
        data: { csrf: $("input[name='csrf']").val() },
        success: function (result) {
            $("#clickname").text(text)
            $("#categorylistid").val(id)
            if (Array.isArray(result.CategoryList)) {
                var stringformat = `<p> ` + result.Count + ` ` + " " + `  ` + result.translate.recordsavailable + ` </p>`
                for (let x of result.CategoryList) {
                    let imagePath = x.ImagePath ? x.ImagePath : "/public/img/space-default-image.png";
                    stringformat += ` <div class="available">
        <div class="category-img">
            <img src="${imagePath}" alt="">
        </div>
        <div class="available-details">
            <p> ${x.CategoryName} </p> 
            <div class="about-available-category">`;
                    for (let y of x.Parent) {
                        stringformat += `   <p>  ${y}  </p>
          <img src="/public/img/caret-right.svg" alt=""> `;
                    }
                    stringformat += `</div>
        </div>
       </div><div class="categories-separate"></div>
        `;
                    // stringformat += ` <div class="available">
                    //     <div class="category-img">
                    //         <img src="`+ x.ImagePath + `" alt="">
                    //     </div>
                    //     <div class="available-details">
                    //         <p> `+ x.CategoryName + ` </p> 
                    //         <div class="about-available-category">`
                    // for (let y of x.Parent) {
                    //     stringformat += `   <p>  ` + y + `  </p>
                    //       <img src="/public/img/caret-right.svg" alt=""> `
                    // }

                    // stringformat += `</div>
                    //     </div>
                    //    </div><div class="categories-separate"></div>
                    //     `
                }

                $(".available-category").html(stringformat)
            } else {
                $(".available-category").html("<p>No records available</p>");
            }
        }

    })

});



// Filter In Sub Category List In MOdal

$("#searchforminclist").keyup(function (event) {
    event.preventDefault()

    var id = $("#categorylistid").val()
    var data = $("#inputcategorylist").val()

    if ($("#searchforminclist").valid()) {
        $.ajax({
            url: "/categories/filtercategory/" + id,
            type: "get",
            datatype: "json",
            data: { "keyword": data, "id": id },
            success: function (result) {

                var stringformat = `<p> ` + result.Count + ` ` + " " + `  ` + result.translate.recordsavailable + ` </p>`
                for (let x of result.CategoryList) {

                    stringformat += ` <div class="available">
                        <div class="category-img">
                            <img src="`+ x.ImagePath + `" alt="">
                        </div>
                        <div class="available-details">
                            <p> `+ x.CategoryName + ` </p> 
                            <div class="about-available-category">`
                    for (let y of x.Parent) {
                        stringformat += `   <p>  ` + y + `  </p>
                          <img src="/public/img/caret-right.svg" alt=""> `
                    }
                    stringformat += `</div>
                        </div>
                       </div><div class="categories-separate"></div>
                     `

                }
                $(".available-category").html(stringformat)

            }

        })
    }

})

// description focus 
const GroupDesc = document.getElementById('category_desc');
const inputGroup = document.querySelectorAll('.input-group');

GroupDesc.addEventListener('focus', () => {

    GroupDesc.closest('.input-group').classList.add('input-group-focused');
});
GroupDesc.addEventListener('blur', () => {
    GroupDesc.closest('.input-group').classList.remove('input-group-focused');
});

$('#save').click(function () {

    jQuery.validator.addMethod("duplicatename", function (value) {

        var result;
        category_id = $("#category_id").val()

        $.ajax({
            url: "/categories/checkcategoryname",
            type: "POST",
            async: false,
            data: { "category_name": value, "category_id": category_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();
            }
        })
        return result.trim() != "true"
    })

    $("#category_form").validate({
        rules: {
            category_name: {
                required: true,
                // categoryname_validator: true
                space: true,
                duplicatename: true,
            },
            category_desc: {
                required: true,
                // categorys_desc: true,
                maxlength: 250,
                space: true,
            }
        },
        messages: {
            category_name: {
                required: "* " + languagedata.Categoryy.catgrpnamevalid,
                space: "* " + languagedata.spacergx,
                duplicatename: "*" + languagedata.Categoryy.namevailderr
            },
            category_desc: {
                required: "* " + languagedata.Categoryy.catgrpdescvalid,
                maxlength: "* " + languagedata.Permission.descriptionchat,
                space: "* " + languagedata.spacergx,
            },
        }
    });

    var formcheck = $("#category_form").valid();
    if (formcheck == true) {
        $('#category_form')[0].submit();
    }
    else {
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })

    }

    return false

})

$('#update-btn').click(function () {

    jQuery.validator.addMethod("duplicatename", function (value) {

        var result;
        category_id = $("#category_id").val()

        $.ajax({
            url: "/categories/checkcategoryname",
            type: "POST",
            async: false,
            data: { "category_name": value, "category_id": category_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();
            }
        })
        return result.trim() != "true"
    })
    $("#category_form").validate({
        rules: {
            category_name: {
                required: true,
                // categoryname_validator: true
                space: true,
                duplicatename: true,
            },
            category_desc: {
                required: true,
                // categorys_desc: true
                maxlength: 250,
                space: true

            }
        },
        messages: {
            category_name: {
                required: "* " + languagedata.Categoryy.catgrpnamevalid,
                space: "* " + languagedata.spacergx,
                duplicatename: " *" + languagedata.Categoryy.namevailderr

            },
            category_desc: {
                required: "* " + languagedata.Categoryy.catgrpdescvalid,
                maxlength: "* " + languagedata.Permission.descriptionchat

            },
        }
    });

    var formcheck = $("#category_form").valid();
    if (formcheck == true) {
        category_id = $("#category_id").val()
        value = $("#category_name").val()

        $.ajax({
            url: "/categories/checkcategoryname",
            type: "POST",
            async: false,
            data: { "category_name": value, "category_id": category_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {

                if (data) {
                    $('#category_form')[0].submit();

                }

            }
        })
    }
    else {
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })

    }

    return false

})

function Validationcheck() {

    if ($('#category_name').hasClass('error')) {
        $('#catname').addClass('input-group-error');
    } else {
        $('#catname').removeClass('input-group-error');
    }

    if ($('#category_desc').hasClass('error')) {
        $('#catdesc').addClass('input-group-error');
    } else {
        $('#catdesc').removeClass('input-group-error');
    }
}
$('.category-cancel').click(function () {
    $('#catname').removeClass('input-group-error')
    $('#catdes').removeClass('input-group-error')
})
$(document).on('click', '#back', function () {

    $('#catname').removeClass('input-group-error')
    $('#catdes').removeClass('input-group-error')
})
// Search return home page
$(document).on('keyup', '.search', function () {

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($('.search').val() === "") {

            window.location.href = "/categories"

        }

    }
})


$(document).on('click', '.selectcheckbox', function () {

    $('#unbulishslt').hide()
    $('#seleccheckboxdelete').removeClass('border-r');

    categorygrbid = $(this).attr('data-id')

  
    if ($(this).prop('checked')) {

        selectedcheckboxarr.push(categorygrbid)

    } else {

        const index = selectedcheckboxarr.indexOf(categorygrbid);

        if (index !== -1) {

            selectedcheckboxarr.splice(index, 1);
        }

        $('#Check').prop('checked', false)

    }


    if (selectedcheckboxarr.length != 0) {

        $('.selected-numbers').removeClass('hidden')

        var items

        if (selectedcheckboxarr.length == 1) {

            items = "Item Selected"
        } else {

            items = languagedata.itemselected
        }

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + items)

        $('#seleccheckboxdelete').removeClass('border-end')

        $('.unbulishslt').html("")

        if (selectedcheckboxarr.length > 1) {

            $('#deselectid').text("Deselect All")

        } else {
            $('#deselectid').text("Deselect")
        }

    } else {

        $('.selected-numbers').addClass('hidden')
    }

    var allChecked = true;

    $('.selectcheckbox').each(function () {

        if (!$(this).prop('checked')) {

            allChecked = false;

            return false;
        }
    });

    $('#Check').prop('checked', allChecked);

})

//ALL CHECKBOX CHECKED FUNCTION//

$(document).on('click', '#Check', function () {

    $('#unbulishslt').hide()
    $('#seleccheckboxdelete').removeClass('border-r');

    selectedcheckboxarr = []

    var isChecked = $(this).prop('checked');

    if (isChecked) {

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selectcheckbox').each(function () {

            categorygrbid = $(this).attr('data-id')

            selectedcheckboxarr.push(categorygrbid)
        })

        $('.selected-numbers').removeClass('hidden')

        var items

        if (selectedcheckboxarr.length == 1) {

            items = "Item Selected"
        } else {

            items = languagedata.itemselected
        }
        $('.checkboxlength').text(selectedcheckboxarr.length + " " + items)

    } else {


        selectedcheckboxarr = []

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selected-numbers').addClass('hidden')
    }

    if (selectedcheckboxarr.length == 0) {

        $('.selected-numbers').addClass('hidden')
    }
})
$(document).on('click', '#seleccheckboxdelete', function () {

    if (selectedcheckboxarr.length > 1) {

        $('.deltitle').text("Delete CategoryGroups?")

        $('#content').text('Are you sure want to delete selected CategoryGroups?')

    } else {

        $('.deltitle').text("Delete CategoryGroup?")

        $('#content').text('Are you sure want to delete selected CategoryGroup?')
    }


    $('#delid').addClass('checkboxdelete')
})
//MULTI SELECT DELETE FUNCTION//
$(document).on('click', '.checkboxdelete', function () {

    var url = window.location.href;

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')


    $('.selected-numbers').hide()
    $.ajax({
        url: '/categories/multiselectcategorygrbdelete',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "categorygrbids": selectedcheckboxarr,
            csrf: $("input[name='csrf']").val(),
            "page": pageno


        },
        success: function (data) {

            if (data.value == true) {

                setCookie("get-toast", "Category Group Deleted Successfully")

                window.location.href = data.url
            } else {

                // setCookie("Alert-msg", "Internal Server Error")

                window.location.href = data.url
                
            }

        }
    })

})

//Deselectall function//

$(document).on('click', '#deselectid', function () {

    $('.selectcheckbox').prop('checked', false)

    $('#Check').prop('checked', false)

    selectedcheckboxarr = []

    $('.selected-numbers').addClass('hidden')

})

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
})

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')
    window.location.href = "/categories"
})

$(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            var value=$(".search").val();
            $(".Closebtn").removeClass("hidden")
            $(".srchBtn-togg").addClass("pointer-events-none")
        } else {
            $(".Closebtn").addClass("hidden")
            $(".srchBtn-togg").removeClass("pointer-events-none")
        }
    });
})
