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

    $("#channelname").text("Select channel")

   
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
    if(channelname != ""){
        $("#channelname").text(channelname)
    }else{
        $("#channelname").text("Select Channel")
    }

    $("#getchannelid").val(channelid)
   
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
    
    if ($(this).prop('checked')) {
       
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

        if (count == 1){
        $("#downloadall").text("Download")
        }else{
            $("#downloadall").text("Download All")

        }
        $("#selectitems").text(arrlenght+" " +'items selected')
        $(".downloadentry").addClass('show')
        $(".downloadentry").removeClass('hidden')
    } else {
        FinalDownloadListArray = jQuery.grep(FinalDownloadListArray, function (value) {

            return value != id;
        });

        DownloadListArray = jQuery.grep(DownloadListArray, function (value) {

                return value != id;

        });
        $("#Check").prop("checked",false)
        $("#exportidarr").val("")
        $(".downloadentry").addClass('hidden')
        $(".downloadentry").removeClass('show')

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
            var count =FinalDownloadListArray.length
       
            $("#downloadall").text("Download All")
         
        $("#selectitems").text(count+" " +'items selected')
           $(this).prop("checked",true)
            $(".downloadentry").addClass('show')
            $(".downloadentry").removeClass('hidden')
        })

     }else{
        
        $(".chlckbox").prop("checked",false)
        $("#exportchannelid").val("")
        // $("#download").prop("disabled",true)
        $(".downloadentry").addClass('hidden')
        $(".downloadentry").removeClass('show')
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

$(document).keydown(function(event) {
    if (event.ctrlKey && event.key === '/') {
        $("#searchentryname").focus().select();
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
$("#downloadentrydata").click(function(){

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

$(document).on("click",'.next-data1',function(){

    var page =$(this).find("a").attr('data-id')

    Entrieslist(channelid,entriesname,statusvalue,username,entriesnamefilter,limit,page)

    FinalDownloadListArray.forEach(element => {
               
        $("#Check"+element).prop("checked",true)
    })

})

$(document).on("click",'.next-data2',function(){

    var page =$(this).find("a").attr('data-id')

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

                    var textclass

                    var classdisabled

                    var classnext

                    var crtpagecount

                    var classcurrent

                    var pgdesign = ""

                    var property =""

                    var npage

                                                         
                    if(data.ChanEntrtlist != null){

                        $(".entries-list").html("")
                        $(".export-pg").html("")
                                        
                        for (let entry of data.ChanEntrtlist) {

                            if ($("#Check").is(':checked')) {
                              
                                property += ` <input type="checkbox" class="hidden peer chlckbox" id="Check`+entry.Id+`" data-id=`+entry.Id +` checked>`
                            //   $("#download").removeAttr("disabled")
                            }else{
                                property += ` <input type="checkbox" class="hidden peer chlckbox" id="Check`+entry.Id+`" data-id=`+entry.Id +`>`
                                // $("#download").prop("disabled",true)


                            }

                                                    
                            if (entry.Status == 0) {
    
                                statusclass = "[#FDF1E4]"
                                textclass="[#FC770F]"
                                changestatus = "Draft"
    
                            } else if (entry.Status == 1) {
    
                                statusclass = "[#E2F7E3]"
                                textclass="[#278E2B]"
                                changestatus = "Published"
    
                            } else {
    
                                statusclass = "[#E2EDFC]"
                                textclass="[#0E56B3]"
                                changestatus = "Unpublished"
                            }
    
                            var entireslist = ` <tr>

                           

                             <td class="px-[16px] py-[12px] pr-0 border-b border-[#EDEDED] align-middle">
                                    <div class="chk-group chk-group-label flex justify-center">
                                     ${property} 
                                    <label for="Check`+entry.Id+`"
                                            class="w-[14px] h-[14px] relative cursor-pointer flex gap-[6px] items-center mb-0 text-[14px] font-normal leading-[1] text-[#262626] tracking-[0.005em]
                                                before:bg-transparent before:w-[14px] before:h-[14px] before:inline-block before:relative before:align-middle before:cursor-pointer before:bg-[url('/public/img/unchecked-box.svg')] before:bg-no-repeat before:bg-contain before:-webkit-appearance-none peer-checked:before:bg-[url('/public/img/checked-box.svg')]  "></label>
                                    </div>
                                </td>

                                 <td class="px-[16px] py-[12px] border-b border-[#EDEDED] align-middle">
                                    <div class="flex gap-[8px] items-center justify-start">
                                        <p class="text-[#262626] font-normal text-xs mb-0 ">`+ entry.Title + ` </p>
                                    </div>
                                </td>


                              <td  class="px-[16px] py-[12px] border-b border-[#EDEDED] text-xs text-[#717171] align-middle">
                                    `+ entry.CreatedDate + `
                                </td>

                                  <td
                                    class="px-[16px] py-[12px] border-b border-[#EDEDED] text-xs text-[#717171] align-middle">
                                    `+ entry.Username + `
                                </td>
                        <td
                                    class="px-[16px] py-[12px] border-b border-[#EDEDED] text-xs text-[#717171] align-middle text-left">
                                    <div class="flex justify-start">
                                        <p
                                            class=" py-[3px] px-1.5 text-${textclass} rounded-sm bg-${statusclass} text-xs font-normal">
                                            ${changestatus}</p>
                                    </div>

                                </td>
                        
            
                        <td>
                            <span id="statuschange" class=""></span>
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
                                    console.log("checkcurrent")
                                    classcurrent = "current"
                                    npage =page 
                                }else{
                                    classcurrent = ""
                                }
                              
                              
                            //  pgdesign += `<li class="next-data"><a  class="${classcurrent}   crtpgno">` +page +` </a></li>`

                                   
                            }
                        }
               
                        if (data.chentrycount >data.limit) {

                              
                            var pagination = "";
                          
                                pagination += ` <li class="next-data2"><a data-id="${data.Pagination.PreviousPage}" ${data.CurrentPage === 1 ? "class='disabled'" : ``}>
                                    <img src="/public/img/carat-right.svg" alt="" />
                                </a></li>`;
                            
                            if (data.CurrentPage > 3) {
                                pagination += `<li class="next-data"><a  class="crtpgno" >1</a></li>`;
                            }
                            if (data.CurrentPage === 5) {
                                pagination += `<li class="next-data"><a   class="crtpgno" >2</a></li>`;
                            }
                            if (data.CurrentPage > 5) {
                                pagination += `<li class="next-data"><a class="crtpgno" >...</a></li>`;
                            }
                            if (data.CurrentPage > 2) {
                                pagination += `<li class="next-data"><a   class="crtpgno">${data.Pagination.TwoBelow}</a></li>`;
                            }
                            if (data.CurrentPage > 1) {
                                pagination += `<li class="next-data"><a  class="crtpgno">${data.Pagination.PreviousPage}</a></li>`;
                            }
                            pagination += `<li class="next-data"><a   ${data.CurrentPage === data.Pagination.TotalPages ? "class='current'" : ``}>${data.CurrentPage}</a></li>`;
                            if (data.CurrentPage < data.Pagination.TotalPages) {
                                pagination += `<li class="next-data"><a class="" >${data.Pagination.NextPage}</a></li>`;
                            }
                           
                            if (data.Pagination.ThreeAfter <= data.Pagination.TotalPages) {
                                pagination += `<li class="next-data"><a class="" >...</a></li>`;
                            }
                            if (data.Pagination.TwoAfter <= data.Pagination.TotalPages) {
                                pagination += `<li class="next-data"><a class="" >${data.Pagination.TotalPages}</a></li>`;
                            }
                            pagination +=`  <li class="next-data1">
                            <a herf="" class="${classnext}" data-id="${data.Pagination.NextPage}">
                                <img src="/public/img/carat-right.svg" alt="" />
                            </a>
                        </li>`
                            paginations=` <ul class="flexx">`+pagination+`</ul> <p>`+data.paginationstartcount+` – `+data.paginationendcount  +` of `+ data.chentrycount+`</p>`
                            $(".export-pg").append(paginations);
                    
                     }else{
                     var pg=`  <p>`+data.paginationstartcount+` – `+data.paginationendcount +` of `+data.chentrycount+`</p>`
                     $(".export-pg").append(pg) 
                     }
                        
                      $('.next-data').each(function(){

                  if ($(this).find('a').text()==data.CurrentPage){

                $(this).find('a').addClass('current')
             }
                      })

                    } else {
                        $(".entries-list").html("")
                        $(".export-pg").html("")

                     if (data.ChanEntrtlist == null && data.filter.Keyword == "" && data.filter.Status == "" && data.filter.Title == "" && data.filter.UserName == "") {

                        var nodata_html = `<tr>
                          <td colspan="6">
                      <div class="max-w-[328px] mx-auto text-center m-[120px_16px]">
                                    <div class="text-center w-fit mx-auto">
                       <img src="/public/img/noData.svg" alt="nodate">
                     </div>
        <h2 class="text-[#262626] text-center text-[18px] font-medium leading-[22.5px] mb-[6px] ">
            `+languagedata.oopsnodata+`</h2>
                                
                            </div>
                        </td>
                     </tr>`
                    //  $("#download").prop("disabled",true)

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
                    //  $("#download").prop("disabled",true)

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

//Deselectall function//

$(document).on('click','#deselectid',function(){

    $('.chlckbox').prop('checked',false)

    $('#Check').prop('checked',false)

    FinalDownloadListArray =[]

    $('#selectitems').addClass('hidden')
    
})