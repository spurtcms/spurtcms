var importfields = []

var Imgfile

var limit = 10

var page = 1

var Mappedfields

var Header

var allData

var fileName

$(document).on('click', '.export-next', function () {

    $("#data-export1").hide()

    $("#data-export2").show()

})

// Select channel // 
$(document).on("click", ".data-drop-list", function () {
    id = $(this).attr("data-id")
    channelname = $(this).attr("data-name")
    selct_id= $("#select_channel").attr("data-id", id)
    $("#select_channel").text(channelname)
    $(this).each(function(){
        if (id == selct_id){
            $(this).addClass("active")
        }else{
            $(this).removeClass("active")
        }
    })
})
// $(document).on('click', '#select_channel', function () {

//     $('#channel-error').hide()

//     var data_id = $('#select_channel').attr("data-id")

//     $('.data-drop-list').each(function () {


//         if (data_id == $(this).find('.chn-drop').attr("data-id")) {

//             $(this).addClass("active")

//             $(this).find('.avaliable-dropdown-itemRight').find('.check-circle').prop('checked', true);

//         } else {
//             $(this).removeClass("active")
//         }
//     })

// })

$(document).on('click', '.im-channel', function () {

    $('#channel-error').hide()

    var C_name = $(this).find('.chn-drop').text()

    var C_id = $(this).find('.chn-drop').attr('data-id')

    $("#select_channel").text(C_name).attr("data-id", C_id)

    $("#chn-name").text(C_name).attr("data-id", C_id)

    $("#channel-field").attr("data-id", C_id)

    $("#select_channel").addClass("select_channel")

    $(".channel-drop").removeClass("show")

})

// xlsx del-icon
$(document).on('click', '.del-xlsx-file', function () {

    $(this).parent().remove()

    $('#xlsx-file').show()
    $("#uploaddrag").show()
    $("#uploadtext").show()

    $("#fileInput").val("")

})

// Xlsx file validation //
$(document).ready(function () {

    $('#fileInput').on('change', function (event) {

        $('#File-error').hide()

        var file = event.target.files[0];

        var filename = $(this).val();

        var ext = filename.split(".").pop().toLowerCase();

        var nameExtract = filename.split("\\").pop()

        fileName = nameExtract.split(".").shift()


        if (file.type === "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet") {

            var reader = new FileReader();

            reader.onload = function (e) {

                var data = e.target.result;

                var workbook = XLSX.read(data, { type: 'binary' });

                var firstSheetName = workbook.SheetNames[0];

                var worksheet = workbook.Sheets[firstSheetName];

                Header = XLSX.utils.sheet_to_json(worksheet, { header: 1 })[0];


                if (Header.length > 0) {

                    allData = XLSX.utils.sheet_to_json(worksheet);

                } else {

                    $('#File-error').text("* No data available in xlsx file").show()

                    $("#fileInput").val("")

                    $("#fileInput").parents('.upload-file').find('.upload-file-btn').show()

                    $(".uploaded-xlsxfile").remove()

                }
            };

            reader.readAsArrayBuffer(file);

        } else {

            $('#File-error').text("* Please select only xlsx file").show()

            $(this).val("")


        }
        if (ext == "xlsx") {

            $(this).parents('.upload-file').find('.upload-file-btn').hide()
            $(this).parents('.upload-file').find('span,p').hide()

            var uploaded_html = `<div class="uploaded-xlsxfile">
                    <img src="/public/img/xlsx.svg" alt="">
                    <p class="para">`+ nameExtract + `</p>
                    <a class="del-xlsx-file"><img src="/public/img/delete.svg" alt=""></a>
                    </div>`

            $("#upload-file").append(uploaded_html)

        } else {

            $(this).val("")

        }
    });
});

// Next-btn //
$(document).on('click', '#nextstep', function () {

    $('#File-error').hide()

    $('#Filezip-error').hide()

    var step = $("#stepnum").val()

    if (step == 1) {

        if ($("#fileInput").val() == "") {

            $('#File-error').text("* Please select your xlsx file").show()

        } if ($("#select_channel").attr("data-id") == "") {

            $('#channel-error').show()

        } if ($("#fileInput").val() != "" & $("#select_channel").attr("data-id") != "") {
            $("#stepnum").val(2)
            $("#step1").hide()
            $("#step2").show()
            $("#previous-btn").text("Previous").addClass("btn-outline").show()

        }

    }
    if (step == 2) {

        if ($("#img-Input").val() == "") {

            $('#Filezip-error').text("* Please select your image file").show()

        }
        if ($("#fileInput").val() != "" & $("#select_channel").attr("data-id") != "" & $("#img-Input").val() != "") {

            $.ajax({
                url: "/channel/channelfields/" + $("#select_channel").attr("data-id"),
                type: "GET",
                dataType: "json",
                data: { "id": $("#select_channel").attr("data-id") },
                success: function (result) {

                    // var dynamicfields = ''

                    // var index = 4

                    // console.log("result.FieldValue", result.FieldValue);

                    // if (result.FieldValue != "") {

                    //     for (let fld of result.FieldValue) {

                    //         dynamicfields = dynamicfields + `<li><button class="para-light" field-id="` + index + `">${fld.FieldName}</button></li>`
                    //         index++;

                    //     }

                    // }

                    for (let [index, field] of Header.entries()) {

                        var html = `   <div class="fields-row field-data">
                    <div class="fields-file">
                        <p class="para">`+ field + `</p>
                    </div>
                    <div class="fields-blog">
                        <div class="dropdown">
                        <input type="hidden" class="input-field" id="input-file`+ index + `">
                            <button class="dropdown-toggle para fields position-relative grid align-items-center" field="`+ index + `" id="select-flds" importfield-id="0" r-field="" data-bs-toggle="dropdown"
                                aria-expanded="false">
                               <span>Select Fields To Import</span>
                                <div class="remove-fld position-absolute border-0 end-0  p-1 bg-white" id="cn-user" style="display:none;"><img src="/public/img/close-1234.svg" alt="clear"></div></button>
                               
                            <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
                                <li>
                                    <form action="">
                                        <div class="filters-search">
                                            <input type="text" placeholder="Select Fields" class="para-light field-srh">
                                            <img src="/public/img/search-filter-img.svg" alt="">
                                        </div>
                                    </form>
                                </li>
    
                                <li>
                                    <button class="para-light ch-fld ch-fld-dd flt" field-id="1" r-field="1" field="`+ index + `">
                                        Title <span class="data-required">*</span>
                                    </button>
                                </li>
    
                                <li>
                                    <button class="para-light ch-fld ch-fld-dd fld"  field-id="2" r-field="2" field="`+ index + `">
                                        Description <span class="data-required">*</span>
                                    </button>
                                </li>
    
                                <li>
                                    <button class="para-light ch-fld ch-fld-dd flm"  field-id="3"  field="`+ index + `">
                                        Image <span class="data-required">*</span>
                                    </button>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>`
                        $(".fields-container").append(html)
                    }

                }
            })

            $(".mapping-opt").show()
            $("#stepnum").val(3)
            $("#step2").hide()
            $("#step3").show()

        }

    }

    var fieldData = []

    if (step == 3) {

        var Title

        var Description

        var Img

        $('.field-data').each(function () {

            var obj = {}

            obj.xlname = $(this).find(".fields-file>p").text()

            obj.fieldname = $(this).find(".fields-blog>.dropdown>.fields>span").text().trim(" ").split(" ")[0]

            if (obj.fieldname == "Title") {

                Title = obj.fieldname

            }
            if (obj.fieldname == "Description") {

                Description = obj.fieldname
            }
            if (obj.fieldname == "Image") {

                Img = obj.fieldname
            }

            if (obj.fieldname != "Select") {

                fieldData.push(obj)

            }


        })

        var Data = []

        console.log("allData", allData);
        console.log("allData", fieldData);

        $.each(allData, function (jsonindex, jsonObj) {

            var matchedObj = {}

            Mappedfields = 0

            $.each(fieldData, function (index, obj) {

                var xlname = obj.xlname;

                var field = obj.fieldname;

                if (jsonObj.hasOwnProperty(xlname)) {

                    var value = jsonObj[xlname];

                    matchedObj[field] = value;

                    Mappedfields++;

                }

            })

            console.log("ewda", Mappedfields, fieldData.length);

            Data.push(matchedObj)


        })


        if (Title == "Title" && Description == "Description" && Img == "Image") {

            var formdata = new FormData();
            formdata.append('image', Imgfile)
            formdata.append('id', $("#select_channel").attr("data-id"))
            formdata.append('csrf', $("input[name='csrf']").val())
            formdata.append('fielddata', JSON.stringify(Data))
            formdata.append('filename', fileName)
            formdata.append('totalfields', Header.length)
            formdata.append('mappedfields', Mappedfields)

            console.log("Fields", Header.length);
            console.log("mappedfields", fieldData.length)
            var unmappedfields = Header.length - fieldData.length
            console.log("mappedfields", unmappedfields)
            // Current time 

            var currentTime = new Date();

            var year = currentTime.getFullYear();
            var month = ('0' + (currentTime.getMonth() + 1)).slice(-2);
            var day = ('0' + currentTime.getDate()).slice(-2);
            var hours = currentTime.getHours();
            var minutes = ('0' + currentTime.getMinutes()).slice(-2);
            var ampm = hours >= 12 ? 'PM' : 'AM';
            hours = hours % 12;
            hours = hours ? hours : 12;
            hours = ('0' + hours).slice(-2);

            var formattedTime = year + ' ' + month + ':' + day + ' ' + hours + ':' + minutes + ' ' + ampm;

            $.ajax({
                url: "/settings/data/importdata",
                type: "post",
                data: formdata,
                processData: false,
                contentType: false,
                success: function (result) {

                    if (result.valid == true) {

                        var htmltxt = `<tr><td>${fileName}</td><td>${formattedTime}</td><td>${allData.length}</td><td>${fieldData.length}</td><td>${unmappedfields}</td></tr>`

                        $("#importfile-dt").append(htmltxt)

                    }
                    // if (result.valid == false) {

                    //     notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>Please select all mandatory fields</span></div>`;

                    //     $(notify_content).insertBefore(".header-rht");

                    //     setTimeout(function () {

                    //         $('.toast-msg').fadeOut('slow', function () {

                    //             $(this).remove();

                    //         });

                    //     }, 5000); // 5000 milliseconds = 5 seconds

                    // } 

                    $(".mapping-opt").hide()
                    $("#stepnum").val(4)
                    $("#next-hide").hide()
                    $("#step3").hide()
                    $("#step4").show()


                }
            })
        } else {

            notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>Please select all mandatory fields</span></div>`;

            $(notify_content).insertBefore(".header-rht");

            setTimeout(function () {

                $('.toast-msg').fadeOut('slow', function () {

                    $(this).remove();

                });

            }, 5000); // 5000 milliseconds = 5 seconds

        }


    }

})
// Select fields //
var field_id

$(document).on('click', '.ch-fld', function () {

    $(".field-srh").val("")

    var fid = $(this).attr("field-id")

    var fieldname = $(this).text()

    field_id = $(this).attr('field-id')

    importfields = []

    $('.input-field').each(function () {

        importfields.push($(this).val())

    })
    const index = importfields.findIndex(id => id === field_id)

    if (index === -1) {

        // $(this).addClass("mapped-field")

        $(this).parents(".fields-blog>.dropdown").find('.fields>span').addClass('select_channel').text(fieldname).attr("importfield-id", field_id)

        $(this).parents(".fields-blog>.dropdown").find('.fields>.remove-fld').show()

        $(this).parents(".fields-blog>.dropdown").find('.input-field').val(field_id)

        if (fid == 1) {

            $(".flt").removeClass("ch-fld-dd")
        }
        if (fid == 2) {

            $(".fld").removeClass("ch-fld-dd")
        }
        if (fid == 3) {

            $(".flm").removeClass("ch-fld-dd")
        }
        importfields = []

        $('.input-field').each(function () {

            importfields.push($(this).val())

        })
    }
})

// reset mapping 
$(document).on("click", "#mp-reset", function () {

    importfields = []

    $(".ch-fld").addClass("ch-fld-dd")

    $('.field-data').each(function () {

        $(this).find(".fields-blog>.dropdown>.fields>span").removeClass('select_channel').text("Select fields to import").attr("importfield-id", '')

        $(this).find(".fields-blog>.dropdown>.fields>.remove-fld").hide()

        $(this).find(".fields-blog>.dropdown>.input-field").val("")

    })
})

// Auto mapping
$(document).on("click", "#auto-mapping", function () {

    importfields = []

    $(".ch-fld").removeClass("ch-fld-dd")

    $('.field-data').each(function (index, _) {

        var This = $(this)

        importfields = []


        $(this).find('.dropdown-menu > li:not(:first) .ch-fld').each(function (val, _) {


            if (index == val) {

                // $(this).addClass("mapped-field")

                $(This).find(".fields-blog>.dropdown>.fields>span").addClass('select_channel').text($(this).text()).attr("importfield-id", $(this).attr('field-id'))
                $(This).find(".fields-blog>.dropdown>.fields>.remove-fld").show()
                $(This).find(".fields-blog>.dropdown").find('.input-field').val($(this).attr('field-id'))


                $('.input-field').each(function () {

                    importfields.push($(this).val())

                })


            }


        })

    })
})


// img-zip 
$('#img-Input').on('change', function (event) {

    $('#Filezip-error').hide()

    Imgfile = this.files[0]

    const zip = new JSZip();

    var filename = $(this).val();

    var ext = filename.split(".").pop().toLowerCase();

    var nameExtract = filename.split("\\").pop()


    if (ext == "zip") {

        zip.loadAsync(Imgfile)
            .then(function (zip) {

                if (Object.keys(zip.files).length === 1) {

                    $('#Filezip-error').text("* Zip file should not be empty").show()

                    $("#img-Input").val("")

                }

                const allowedExtensions = ['jpeg', 'jpg', 'png'];

                zip.forEach(function (_, zipEntry) {

                    if (!zipEntry.dir) {

                        const extension = zipEntry.name.split('.').pop().toLowerCase();

                        if (allowedExtensions.indexOf(extension) === -1) {

                            $("#img-Input").val("")

                            $('#Filezip-error').text("* Zip file should contain only jpg/jpeg/png format ").show()

                        } else {

                            console.log('Valid file format:', zipEntry.name);
                        }
                    }
                });
            })
        $(this).parents('.upload-file').find('.upload-file-btn').hide()
        $(this).parents('.upload-file').find('span,p').hide()
        var uploaded_html = `<div class="uploaded-xlsxfile">
            
            <p class="para">`+ nameExtract + `</p>
            <a class="del-zip-file"><img src="/public/img/delete-icon.svg" alt=""></a>
            </div>`

        $("#upload-zipfile").append(uploaded_html)

    } else {

        $(this).val("")

        $('#Filezip-error').text("* Please select only image zip file").show()

    }

})

// xlsx del-icon
$(document).on('click', '.del-zip-file', function () {

    $(this).parent().remove()

    $('#zip-file').show()

    $("#img-Input").val("")

})

// previous 
$(document).on('click', '#previous-btn', function () {

    var step = $("#stepnum").val()

    if (step == 2) {

        $("#stepnum").val(1)
        $("#step2").hide()
        $("#step1").show()
        $("#previous-btn").text("Cancel").removeClass("btn-outline").hide()

    }
    if (step == 3) {
        $(".mapping-opt").hide()
        $("#stepnum").val(2)
        $("#step3").hide()
        $("#step2").show()
        $("#previous-btn").text("Previous").addClass("btn-outline")
        $(".field-data").remove()
    }
})

$(document).on("click", "#err-xldata", function (event) {

    event.preventDefault()

    window.location.href = "/settings/data/err-xlsx-download"

    if (document.cookie.includes("redirect_after_download=true")) {

        document.cookie = "redirect_after_download=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";

        setTimeout(() => {

            window.location.href = "/settings/data/"

        }, 1000)

    }


})

// channel fields search

$(document).on("keyup", ".field-srh", function () {

    var value = $(this).val().toLowerCase();


    $(".ch-fld").filter(function () {

        $(this).parent('li').toggle($(this).text().toLowerCase().indexOf(value) > -1)

    });

})

$(document).on("click", "#select-flds", function () {

    $("li").show();

})

$(".import-map-list").click(function () {

    var chlid

    var C_id = $(this).find('#chn-name').attr('data-id')

    $(".import-map").each(function () {
        $(this).removeClass("active")
        $(this).find(".check-circle").prop("checked", false)
        chlid = $(this).attr("data-id")

        if (chlid == C_id) {
            $(this).addClass("active")
            $(this).find(".check-circle").prop("checked", true)
        }
    })

})
$(document).on("click", ".import-map", function () {

    channelid = $(this).attr("data-id")
    channelname = $(this).attr("data-name")

    $("#select_channel").text(channelname).attr("data-id", channelid)
    $("#chn-name").text(channelname).attr("data-id", channelid)

})

$(document).on("click", ".remove-fld", function () {

    var fid = $(this).parents(".fields-blog>.dropdown").find('.input-field').val()

    $(this).parents(".fields-blog>.dropdown").find('.fields>span').removeClass('select_channel').text("Select Fields To Import").attr("importfield-id", '')

    $(this).hide()

    $(this).parents(".fields-blog>.dropdown").find('.input-field').val("")

    if (fid == 1) {

        $(".flt").addClass("ch-fld-dd")
    }
    if (fid == 2) {

        $(".fld").addClass("ch-fld-dd")
    }
    if (fid == 3) {

        $(".flm").addClass("ch-fld-dd")
    }
    importfields = []

    $('.input-field').each(function () {

        importfields.push($(this).val())

    })
})
