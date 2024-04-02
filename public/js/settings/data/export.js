$(document).ready(function(){

    if ($('#statusid').val() == '') {
        $('#triggerId').val('Status');
    }

    if ($('#statusid').val() == 'Draft') {
        $('#triggerId').val('Draft');
    }
    if ($('#statusid').val() == 'Published') {
        $('#triggerId').val('Published');
    }
    if ($('#statusid').val() == 'Unpublished') {
        $('#triggerId').val('Unpublished');
    }

   
})

// Get Channel id and name 
var channelid

var channelname

var limit = 10

var page = 1

$(".chllist-dropdown").click(function () {

    channelid = $(this).attr("data-id")
    channelname = $(this).attr("data-name") 
    
    var crtchannelid = $("#channelid").val(channelid)
    $("#channelname").text(channelname)
    $("#channel-dropdown").css("color","#112D55")
   
    if (crtchannelid != undefined){
        $('#dropdown-check').removeClass('input-group-error');
        $("#channel-dropdown").css("border","1px solid #EDEDED")
        $("#channelid-error").css("display", "none")
    }else{
        $('#dropdown-check').addClass('input-group-error');
        $("#channel-dropdown").css("border", "1px solid #F26674")
        $("#channelid-error").show()
       
    }

    Entrieslist(channelid,"","","","",limit,page)

  
})

$(".channel-select-list").click(function(){
   
    var chlid 

    $(".chllist-dropdown").each(function() {
        $(this).removeClass("active")
        $(this).find(".check-circle").prop("checked",false)
        chlid = $(this).attr("data-id")
       
       if(chlid == channelid){
        $(this).addClass("active")
        $(this).find(".check-circle").prop("checked",true)
       }
    })
      
})


// entries list
$("#export").click(function () {

    if (channelid == undefined) {
        
        $('#dropdown-check').addClass('input-group-error');
        $("#channel-dropdown").css("border", "1px solid #F26674")
        $("#channelid-error").show()
    } else {
         $("#export").addClass("export-next")
         $("#getchannelid").val(channelid) 
          
    }

   Entrieslist(channelid,"","","","",limit,page)
   
})

var DownloadListArray =[];

var FinalDownloadListArray = [];

/*ckbox*/
$(document).on('click', '.chlckbox', function () {

    DownloadListArray=[]

    var id = $(this).attr("data-id")

    var count = $("#entrycount").val()

    if ($(this).is(':checked')) {

        DownloadListArray.push(id)

        DownloadListArray.forEach(element => {

            if (!FinalDownloadListArray.includes(element)) {

                FinalDownloadListArray.push(element);
            }
        })

        var arrlenght = FinalDownloadListArray.length


        if(count == arrlenght){
            $("#Check").prop("checked",true)
        }
        $("#download").removeAttr("disabled")
    } else {
        FinalDownloadListArray = jQuery.grep(FinalDownloadListArray, function (value) {

            return value != id;
        });

        DownloadListArray = jQuery.grep(DownloadListArray, function (value) {

                return value != id;

        });
        $("#Check").prop("checked",false)
        $("#exportidarr").val("")
        $("#download").prop("disabled",true)

    }

    $("#exportidarr").val(FinalDownloadListArray)
    console.log("checking array", FinalDownloadListArray);


})

var channelid

// export all data

$("#Check").click(function(){

    DownloadListArray=[]

    if ($(this).is(':checked')) {

        $(".chlckbox").each(function(){

          var checkboxid = $(this).attr("data-id")

           channelid = $("#channelid").val()

            DownloadListArray.push(checkboxid)

            DownloadListArray.forEach (element =>{

                if (!FinalDownloadListArray.includes(element)) {

                    FinalDownloadListArray.push(element);
                }
            })
        
            $(this).prop("checked",true)
            $("#download").removeAttr("disabled")

        })

     }else{
        
        $(".chlckbox").prop("checked",false)
        $("#exportchannelid").val("")
        $("#download").prop("disabled",true)
    }
    $("#exportchannelid").val(channelid)

    console.log("downl arra",FinalDownloadListArray);
})

// Search Functionality
var entriesname 

document.addEventListener('keydown', function(event) {
    if (event.key === "Enter") {
        entriesname = $("#searchentryname").val()
        Entrieslist(channelid ,entriesname,"","","",limit,page)
    }
});

/*search redirect home page */

$(document).on('keyup', '#searchentryname', function () {
   
    if ($('.search').val() === "") {
        $('.noData-foundWrapper').remove()
        Entrieslist(channelid,"","","","",limit,page)
    }
})

// empty the check box
$("#download").click(function(){

    $("#Check").prop("checked",false)

    FinalDownloadListArray.forEach(element =>{

        $("#Check"+element).prop("checked",false)
        // $(this).prop("disabled",true)


    })

    FinalDownloadListArray =[]

})

// Dropdown status filter

var statusvalue

var entriesnamefilter

var username

$(document).on('click', '.statuss', function () {

    statusvalue = $(this).text()
    console.log("statusvalue",statusvalue);
    $('#triggerId').val(statusvalue)
    $('#statusid').val(statusvalue)
    $('.filter-dropdown-menu').addClass('show').css({
        position: 'absolute',
        inset: '0px 0px auto auto',
        margin: '0px',
        transform: 'translate(0px, 34px)'
    });


})
// Cancel button in filter

$(document).on('click', '#filtercancel', function () {

    $('#triggerId').val("Status");
    $("#title").val("")
    $("#statusid").val("")
    $("#username").val("")
    Entrieslist(channelid,"","","","",limit,page)
});

// Apply Filter functionality

$("#applyfilter").click(function(){

    entriesnamefilter = $("#title").val()

    username = $("#username").val()

    console.log("values",entriesnamefilter,username,statusvalue);

    if(statusvalue != undefined || entriesnamefilter != undefined || username != undefined){
        Entrieslist(channelid,"",statusvalue,username,entriesnamefilter,limit,page)
        
    }else{
      
        console.log("error");
    }

    $("#search-dropdown").removeClass("show")

    $('.selected-list').show()
  
})


// pagantion functionality

$(document).on("click",".next-data", function(){
   
    var page =$(this).find("a").text().trim(" ")

    Entrieslist(channelid,entriesname,statusvalue,username,entriesnamefilter,limit,page)

    FinalDownloadListArray.forEach(element => {
               
        $("#Check"+element).prop("checked",true)
    })

    
})

// entries list
function Entrieslist(channelid,entriesname,statusvalue,username,entriesnamefilter,limit,page){

    
    if (channelid != undefined) {

        $.ajax({
            url: "/settings/data/entrieslist/" + channelid,
            type: "GET",
            async: false,
            data:{"id":channelid,"keyword":entriesname,"Status":statusvalue,"Username":username,"Title":entriesnamefilter,"limit":limit,"page":page},
            datatype: "json",
            caches: false,
            success: function (data) {

                console.log(data);
                          
                if (data.hasOwnProperty('ChanEntrtlist')) {

                    var arrlenght = data.chentrycount
                    $("#entrycount").val(arrlenght)
                    var statusclass

                    var changestatus

                    var classdisabled

                    var classnext

                    var crtpagecount

                    var classcurrent

                    var pgdesign = ""

                    var property =""

                                                         
                    if(data.ChanEntrtlist != null){

                        $(".entries-list").html("")
                        $(".export-pg").html("")
                                        
                        for (let entry of data.ChanEntrtlist) {

                            if ($("#Check").is(':checked')) {
                                property += ` <input type="checkbox" class="chlckbox" id="Check`+entry.Id+`" data-id=`+entry.Id +` checked>`
                              $("#download").removeAttr("disabled")
                            }else{
                                property += ` <input type="checkbox" class="chlckbox" id="Check`+entry.Id+`" data-id=`+entry.Id +`>`
                                $("#download").prop("disabled",true)


                            }

                                                    
                            if (entry.Status == 0) {
    
                                statusclass = "status draft"
                                changestatus = "Draft"
    
                            } else if (entry.Status == 1) {
    
                                statusclass = "status published"
                                changestatus = "Published"
    
                            } else {
    
                                statusclass = "status unpublished"
                                changestatus = "Unpublished"
                            }
    
                            var entireslist = ` <tr>

                            <td>
                        <div class="check-id">
                        <div class="chk-group">
                        ${property}
                        <label for="Check`+entry.Id+`"></label>
                    </div>
                            <p class="para">`+ entry.Cno + `</p>
                        </div>
                     </td>

                        <td>`+ entry.Title + ` </td>
                        <td> `+ entry.CreatedDate + ` </td>
                        <td>`+ entry.Username + ` </td>
            
                        <td>
                            <span id="statuschange" class="${statusclass}">${changestatus}</span>
                        </td>

                        <td></td>
                                    
                       </tr>`
    
                            $(".entries-list").append(entireslist)
    
                        }

                        crtpagecount =data.CurrentPage

                        if(data.CurrentPage == 1){
                            classdisabled = "disabled"
                        }

                        if(data.CurrentPage == data.PageCount){
                            classnext = "disabled"
                        }else{
                            classnext = "next"
                        }

                        if(data.Page !=null){
                            for (let page of data.Page) {
                                
                                if(page == crtpagecount){
                                    classcurrent = "current"
                                }else{
                                    classcurrent = ""
                                }

                              
                             pgdesign += `<li class="next-data"><a  class="${classcurrent}   crtpgno">` +page +` </a></li>`

                                   
                            }
                        }
                    

                       if(data.Page != null){
                        var paganation =` <ul class="flexx">
                        <li>
                            <a herf="" class="${classdisabled}">
                                <img src="/public/img/carat-right.svg" alt="" />
                            </a>
                        </li>
                        ${pgdesign}

                        <li>
                            <a herf="" class="${classnext}" >
                                <img src="/public/img/carat-right.svg" alt="" />
                            </a
                        </li>
                     </ul>

                       <p>`+data.paginationstartcount+` – `+data.paginationendcount  +` of `+ data.chentrycount+`</p>

                     `

                     $(".export-pg").append(paganation) 
                     }else{
                     var pg=`  <p>`+data.paginationstartcount+` – `+data.paginationendcount +` of `+data.chentrycount+`</p>`
                     $(".export-pg").append(pg) 
                     }
                        


                    } else {
                        $(".entries-list").html("")
                        $(".export-pg").html("")

                     if (data.ChanEntrtlist == null && data.filter.Keyword == "" && data.filter.Status == "" && data.filter.Title == "" && data.filter.UserName == "") {

                        var nodata_html = `<tr class="no-dataHvr">
                        <td style="text-align: center;border-bottom: none;" colspan="6">
                            <div class="noData-foundWrapper">

                                <div class="empty-folder">
                                    <img src="/public/img/folder-sh.svg" alt="">
                                    <img src="/public/img/shadow.svg" alt="">
                                </div>
                                <h1 class="heading">`+languagedata.oopsnodata+`</h1>
                            </div>
                        </td>
                     </tr>`

                     $(".entries-list").append(nodata_html)

                    } else if (data.filter.Keyword != "" || data.filter.Status != "" || data.filter.Title != "" || data.filter.UserName != "" && data.ChanEntrtlist == null  ){

                        var nodata_html = `<tr class="no-dataHvr">
                        <td style="text-align:center;border-bottom: none;"colspan="6">
                            <div class="noData-foundWrapper">

                                <div class="empty-folder">
                                    <img src="/public/img/nodatafilter.svg" alt="">
                                </div>
                                <h1 class="heading">`+languagedata.oopsnodata+`</h1>
                            </div>
                        </td>
                     </tr>`

                     $(".entries-list").append(nodata_html)

                    } else {

                        $('.noData-foundWrapper').remove()
                    }

                 }
                                

                }
                
            }
        })
    } 

}


$(document).on('click','#removecategory',function(){


    $('#triggerId').val("Status");
    $("#title").val("")
    $("#statusid").val("")
    $("#username").val("")
    Entrieslist(channelid,"","","","",limit,page)
    $(".selected-list").hide()

})

// search key value empty in dropdown search

// $(".channel-select-list").on("click",function(){

//     var keyword = $("#searchentrylists").val()
//     $(".chl-dropdown").each(function (index, element) {
//         if(keyword != ""){
//         $("#searchentrylists").val("")
//         $("#entrynodata").hide()
//         $(element).show()
//       }else{
//         $("#searchentrylists").val()
//         $("#entrynodata").hide()
//         $(element).show()
    
//       }
//     })
   
//   })


//   // dropdown filter input box search
// $("#searchentrylists").keyup(function () {
//     var keyword = $(this).val().trim().toLowerCase()
//     $(".chl-dropdown").each(function (index, element) {
//         var title = $(element).attr("data-name").toLowerCase()

//         if (title.includes(keyword)) {
//             $(element).show()
//             $("#entrynodata").hide()

//         } else {
            
//             $(element).hide()
//             if($(".chl-dropdown:visible").length==0){
//                 $("#entrynodata").show()
//             }
           
//         }
//     })

// })