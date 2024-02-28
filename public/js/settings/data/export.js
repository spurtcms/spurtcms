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

$(".chl-dropdown").click(function () {
    channelid = $(this).attr("data-id")
    var channelname = $(this).attr("data-name") 

    var crtchannelid = $("#channelid").val(channelid)
    $("#chennalname").text(channelname)


    if (crtchannelid != undefined){
        $('#dropdown-check').removeClass('input-group-error');
        $("#channel-dropdown").css("border","1px solid #EDEDED")
        $("#channelid-error").css("display", "none")
    }else{
        $('#dropdown-check').addClass('input-group-error');
        $("#channel-dropdown").css("border", "1px solid #F26674")
        $("#channelid-error").show()
    }

   
})

// entries list
$("#export").click(function () {

    if (channelid == undefined) {
        
        $('#dropdown-check').addClass('input-group-error');
        $("#channel-dropdown").css("border", "1px solid #F26674")
        $("#channelid-error").show()
    } else {
         $("#export").addClass("export-next")

    }

   Entrieslist(channelid)
   
})

var DownloadListArray =[];

/*ckbox*/
$(document).on('click', '.chlckbox', function () {

    var id = $(this).attr("data-id")

    console.log("id",id);

    if ($(this).is(':checked')) {

        DownloadListArray.push(id)

        $("#exportidarr").val(DownloadListArray)

    } else {

        // DeleteListArray = jQuery.grep(DeleteListArray, function (value) {

        //     return value != name;

        // });
    }

    console.log("checking array", DownloadListArray);


})

$("#Check").click(function(){

    $(".chlckbox").prop("checked",true)

})
var  items

var item
// download functionlity

$("#download").click(function(){

    if (DownloadListArray.length == 0) {

        var notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>Please select one entries</span></div>';

            $(notify_content).insertBefore(".header-rht");

    }else{

            // $.ajax({
            //             url:"/settings/data/export",
            //             type:"GET",
            //             data:{id:DownloadListArray} ,
            //             caches: false,
            //             success:function(){
            //              console.log("download successfully");
            //             },
            //             error: function(){
            //                 console.log("Error on response data");
            //             } 
            //  });
         
    }

    setTimeout(function () {

        $('.toast-msg').fadeOut('slow', function () {

            $(this).remove();

        });

    }, 5000);
})


// Search Functionality
var entriesname 

document.addEventListener('keydown', function(event) {
    if (event.key === "Enter") {
        entriesname = $("#searchentryname").val()
        Entrieslist(channelid ,entriesname)
    }
});

/*search redirect home page */

$(document).on('keyup', '#searchentryname', function () {
    if ($('.search').val() === "") {
        Entrieslist(channelid)
    }
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

    $('#triggerId').val('');
    Entrieslist(channelid)
});

// Apply Filter functionality

$("#applyfilter").click(function(){

    entriesnamefilter = $("#title").val()

    username = $("#username").val()

   Entrieslist(channelid,entriesname,statusvalue,username,entriesnamefilter)

    // $("#search-dropdown").hide()
   
})

 // filter dropdown show
// $(".filter-dropdown-menu").click(function(){
//     $(this).show()
// })


function Entrieslist(channelid,entriesname,statusvalue,username,entriesnamefilter){

    
    if (channelid != undefined) {

        $.ajax({
            url: "/settings/data/entrieslist/" + channelid,
            type: "GET",
            async: false,
            data:{"id":channelid,"keyword":entriesname,"Status":statusvalue,"Username":username,"Title":entriesnamefilter},
            datatype: "json",
            caches: false,
            success: function (data) {

                console.log(data);
                          

                if (data.hasOwnProperty('ChanEntrtlist')) {

                    var statusclass

                    var changestatus

                    var classdisabled

                    var classnext

                    var crtpagecount

                    var classcurrent

                                      
                    if(data.ChanEntrtlist != null){

                                        
                        $(".entries-list").html("")
                        $(".export-pg").html("")

                        for (let entry of data.ChanEntrtlist) {

                                                     
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
                                <input type="checkbox" class="chlckbox" id="Check`+entry.Id+`" data-id=`+entry.Id +`>
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
                            classdisabled = "disabled"
                        }else{
                            classnext = "next"
                        }

                        if(data.Page !=null){
                            for (let page of Page) {

                                if(page == crtpagecount){
                                    classcurrent = "current"
                                }
    
                            }
                        }

                      

               if(data.Page != null){
                var paganation =`     <ul class="flexx">
                        <li>
                            <a href="?page=`+data.Previous +`" class="${classdisabled}">
                                <img src="/public/img/carat-right.svg" alt="" />
                            </a>
                        </li>
                         
                         <li><a href="?page=`+page +`" class="${classcurrent}  ">` +page +` </a></li>
                        
                        <li>
                            <a href="?page=`+data.Next +`"  class="${classdisabled}"  class="${classnext}"
                                >
                                <img src="/public/img/carat-right.svg" alt="" />
                            </a>
                        </li>
                    </ul>`
                    $(".export-pg").append(paganation) 
                }else{
                  var pg=`  <p>`+data.paginationstartcount+` – `+data.paginationendcount +`of `+data.chentrycount+`</p>`
                   $(".export-pg").append(pg) 
                }
                        


                    } else{
                    if ($('.entries-list:visible').length == 0) {

                        var nodata_html = `<tr>
                        <td style="text-align: center;" colspan="6">
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

                      } else {

                        $('.noData-foundWrapper').remove()
                      }
                }

                   

                }
                

            }
        })
    } 

}