var importfields = []
$(document).on('click', '.export-next', function () {
    $("#data-export1").hide()
    $("#data-export2").show()
})

$(document).on('click', '#stp-head1', function () {
    $("#stepnum").val(1)
    $("#step2").hide()
    $("#step3").hide()
    $("#step1").show()
    $("#next-hide").show()
})
$(document).on('click', '#stp-head2', function () {
    $("#stepnum").val(2)
    $("#step1").hide()
    $("#step3").hide()
    $("#step2").show()
    $("#next-hide").show()
})
$(document).on('click', '#stp-head3', function () {
    $("#stepnum").val(3)
    $("#step1").hide()
    $("#step2").hide()
    $("#next-hide").hide()
    $("#step3").show()
})
// Select channel // 
$(document).on('click', '#select_channel', function () {

    $('#channel-error').hide()

    var data_id = $('#select_channel').attr("data-id")

    $('.data-drop-list').each(function () {


        if (data_id == $(this).find('.chn-drop').attr("data-id")) {

            $(this).addClass("active")

            $(this).find('.avaliable-dropdown-itemRight').find('.check-circle').prop('checked', true);

        } else {
            $(this).removeClass("active")
        }
    })

})
$(document).on('click', '.data-drop-list', function () {

    $('#channel-error').hide()

    var C_name = $(this).find('.chn-drop').text()

    var C_id = $(this).find('.chn-drop').attr('data-id')

    $("#select_channel").text(C_name).attr("data-id", C_id)

    // $(".data-head-txt").find('.heading').text(C_name).attr("data-id", C_id)

    $("#select_channel").addClass("select_channel")

    $(".channel-drop").removeClass("show")

    console.log("chan", C_name, C_id);
})

// File validation //
var Header
var allData
$(document).ready(function () {
    $('#fileInput').on('change', function (event) {
        $('#File-ms-error').hide()
        $('#File-em-error').hide()
        $('#File-error').hide()
        var file = event.target.files[0];

        if (file.type === "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet") {
            var reader = new FileReader();

            reader.onload = function (e) {
                var data = e.target.result;
                console.log("data", data);
                var workbook = XLSX.read(data, { type: 'binary' });
                console.log("workbook", workbook);
                var firstSheetName = workbook.SheetNames[0];
                console.log("firstSheetName", firstSheetName);
                var worksheet = workbook.Sheets[firstSheetName];
                Header = XLSX.utils.sheet_to_json(worksheet, { header: 1 })[0];

                if (Header.length > 0) {
                    console.log('First Row:', Header);

                    allData = XLSX.utils.sheet_to_json(worksheet);
                    console.log('All Data:', allData);
                } else {
                    $('#File-em-error').show()
                    console.log('File is empty.');
                }
            };

            reader.readAsArrayBuffer(file);
        } else {
            $('#File-ms-error').show()
            console.log('File is not of type XLSX.');
        }
    });
});

// Next-btn //
$(document).on('click', '#nextstep', function () {
    var step = $("#stepnum").val()
    if (step == 1) {
        if ($("#fileInput").val() == "") {

            $('#File-error').show()

        } if ($("#select_channel").attr("data-id") == "") {

            $('#channel-error').show()

        } if ($("#fileInput").val() != "" & $("#select_channel").attr("data-id") != "") {
            $.ajax({
                url: "/channel/channelfields/" + $("#select_channel").attr("data-id"),
                type: "GET",
                dataType: "json",
                data: { "id": $("#select_channel").attr("data-id") },
                success: function (result) {
                    console.log("data", result.FieldValue);

                    var dynamicfields = ''
                    var index = 4
                    for (let fld of result.FieldValue) {
                        dynamicfields = dynamicfields + `<li><button class="para-light" field-id="` + index + `">${fld.FieldName}</button></li>`
                        index++;
                    }
                    for (let [index, field] of Header.entries()) {

                        var html = `   <div class="fields-row">
                        <div class="fields-file">
                            <p class="para">`+ field + `</p>
                        </div>
                        <div class="fields-blog">
                            <div class="dropdown">
                                <button class="dropdown-toggle para fields" field="`+ index + `"  importfield-id="0" data-bs-toggle="dropdown"
                                    aria-expanded="false">
                                    Select Fields To Import
                                </button>
                                <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
                                    <li>
                                        <form action="">
                                            <div class="filters-search">
                                                <input type="text" placeholder="Select Fields" class="para-light">
                                                <img src="/public/img/search-filter-img.svg" alt="">
                                            </div>
                                        </form>
                                    </li>
        
                                    <li>
                                        <button class="para-light"  field-id="1" field="`+ index + `">
                                            Title <span class="data-required">*</span>
                                        </button>
                                    </li>
        
                                    <li>
                                        <button class="para-light"  field-id="2" field="`+ index + `">
                                            Description
                                        </button>
                                    </li>
        
                                    <li>
                                        <button class="para-light"  field-id="3" field="`+ index + `">
                                            Image
                                        </button>
                                    </li>
                                    ${dynamicfields}
                                </ul>
                            </div>
                        </div>
                    </div>`

                        $(".fields-container").append(html)
                    }
                }
            })


            $("#stepnum").val(2)
            $("#step1").hide()
            $("#step2").show()
        }

    }
    if (step == 2) {
        $("#stepnum").val(3)
        $("#next-hide").hide()
        $("#step2").hide()
        $("#step3").show()
    }

})

// Select fields //

$(document).on('click', '.para-light', function () {

    var fieldname = $(this).text()

    var field_id = $(this).attr('field-id')

    var field_Index = $(this).attr('field')

    var This = $(this).parents(".fields-blog>.dropdown")

    const index = importfields.findIndex(obj => obj.id === field_id)

    if (index === -1) {

        var field_Obj = {}

        field_Obj.id = field_id

        field_Obj.fieldIndex = field_Index

        $(this).parents(".fields-blog>.dropdown").find('.fields').addClass('select_channel').text(fieldname).attr("importfield-id", field_id)

        importfields.push(field_Obj)


    } else {
        for (let x of importfields) {

            $('.fields').each(function () {

                console.log("field", (x.id == field_id) && ($(this).attr("importfield-id") != x.id));

                console.log("val", $(This).find('.fields').attr("importfield-id"));
                
                if ((x.id == field_id) && ($(this).attr("importfield-id") == field_id) && ($(This).find('.fields').attr("importfield-id") != field_id)) {

                    importfields.splice(x, 1)

                }
            })
        }



    }
    console.log("fields", importfields);

    // parentfield = $(this).parents(".fields-blog>.dropdown").find('.fields').attr('field'

    // $('.fields').each(function () {

    //     if($(this)==parent){

    //         $(this).addClass('select_channel')

    //         $(this).text(fieldname)

    //         $(this).attr("field-id", field_id)
    //     }
    // })

    // $('.fields').each(function () {

    // for (let fieldid of importfields){

    //     if($(this).attr('importfield-id')==fieldid){

    //     $(this).text(fieldname)
    // }
    //     }

    // })


})