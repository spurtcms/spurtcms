/**Variables */
var languagedata;

let editor;

let ckeditor1;

let collapsecount = 1

let newPage = 0

let OrderIndex = 1

let PageId = 1

let SubPageId = 1

let GroupId = 1

let newpagesarr = []

let newGrouparr = []

let Subpagearr = []

let NewSubObj = []

let overallarray = []

/**delete array*/
let deletePagearr = []

let deletegrouparr = []

let deletesubarr = []

/** */
// $(document).ready(async function () {

    
// })


$(document).ready( async function () {

    var languagecode = $('.language-group>button').attr('data-code')

    await  $.getJSON("/locales/"+languagecode+".json", function (data) {
        
        languagedata = data
    })


    CKEDITORS()

   $.ajax({
        type: "get",
        url: "/spaces/page/pagedata",
        dataType: 'json',
        data: {
            sid: $("#spid").val(),
        },
        cache: false,
        success: function (result) {

        }
    })


    if (newpagesarr.length >= 0) {

        var pagename = languagedata.Spaces.introduction

        var pagedata = AddPageString(pagename, PageId,languagedata)

        $('.newpages').append(pagedata)

        var pageobj = PageObjectCreate();

        pageobj.Name = pagename

        newpagesarr.push(pageobj)

        PageId++;

        OrderIndex++;

    }

    $('.viewpage').each(function () {

        $(this).trigger('click');

        return;
    })


})

function CKEDITORS() {

    CKEDITOR.ClassicEditor.create(document.getElementById("neweditor"), {
        toolbar: {
            items: ['heading', 'bold', 'italic','blockQuote', 'imageUpload',  'numberedList', 'bulletedList', 'todoList','horizontalLine','link','code'],
            shouldNotGroupWhenFull: true
        },
        // toolbar:{
        //     items:['alignment:left', 'alignment:right', 'alignment:center', 'alignment:justify', 'alignment', 'bold', 'italic', 'underline', 'strikethrough', 'superscript', 'subscript', 'code', 'blockQuote', 'codeBlock', 'selectAll', 'undo', 'redo', 'heading', 'imageTextAlternative', 'toggleImageCaption', 'imageStyle:inline', 'imageStyle:alignLeft', 'imageStyle:alignRight', 'imageStyle:alignCenter', 'imageStyle:alignBlockLeft', 'imageStyle:alignBlockRight', 'imageStyle:block', 'imageStyle:side', 'imageStyle:wrapText', 'imageStyle:breakText', 'imageUpload', 'insertImage', 'imageInsert', 'indent', 'outdent', 'textPartLanguage', 'link', 'linkImage', 'numberedList', 'bulletedList', 'todoList', 'mediaEmbed', 'findAndReplace', 'fontBackgroundColor', 'fontColor', 'fontFamily', 'fontSize', 'highlight:yellowMarker', 'highlight:greenMarker', 'highlight:pinkMarker', 'highlight:blueMarker', 'highlight:redPen', 'highlight:greenPen', 'removeHighlight', 'highlight', , 'htmlEmbed', 'pageBreak', 'removeFormat', 'restrictedEditingException', 'showBlocks', 'style', 'specialCharacters', 'insertTable', 'tableColumn', 'tableRow', 'mergeTableCells', 'tableCellProperties', 'tableProperties', 'toggleTableCaption', 'sourceEditing']
        // },
        list: {
            properties: {
                styles: true,
                startIndex: true,
                reversed: true
            }
        },
        heading: {
            options: [
                { model: 'paragraph', title: 'Paragraph', class: 'ck-heading_paragraph' },
                { model: 'heading1', view: 'h1', title: 'Heading 1', class: 'ck-heading_heading1' },
                { model: 'heading2', view: 'h2', title: 'Heading 2', class: 'ck-heading_heading2' },
                { model: 'heading3', view: 'h3', title: 'Heading 3', class: 'ck-heading_heading3' },
                { model: 'heading4', view: 'h4', title: 'Heading 4', class: 'ck-heading_heading4' },
                { model: 'heading5', view: 'h5', title: 'Heading 5', class: 'ck-heading_heading5' },
                { model: 'heading6', view: 'h6', title: 'Heading 6', class: 'ck-heading_heading6' }
            ]
        },
        placeholder: languagedata.Spaces.pllckeditor,
        fontFamily: {
            options: [
                'default',
                'Arial, Helvetica, sans-serif',
                'Courier New, Courier, monospace',
                'Georgia, serif',
                'Lucida Sans Unicode, Lucida Grande, sans-serif',
                'Tahoma, Geneva, sans-serif',
                'Times New Roman, Times, serif',
                'Trebuchet MS, Helvetica, sans-serif',
                'Verdana, Geneva, sans-serif'
            ],
            supportAllValues: true
        },
        fontSize: {
            options: [10, 12, 14, 'default', 18, 20, 22],
            supportAllValues: true
        },
        htmlSupport: {
            allow: [
                {
                    name: /.*/,
                    attributes: true,
                    classes: true,
                    styles: true
                }
            ]
        },
        htmlEmbed: {
            showPreviews: true
        },
        link: {
            decorators: {
                addTargetToExternalLinks: true,
                defaultProtocol: 'https://',
                toggleDownloadable: {
                    mode: 'manual',
                    label: 'Downloadable',
                    attributes: {
                        download: 'file'
                    }
                }
            }
        },
        mention: {
            feeds: [
                {
                    marker: '@',
                    feed: [
                        '@apple', '@bears', '@brownie', '@cake', '@cake', '@candy', '@canes', '@chocolate', '@cookie', '@cotton', '@cream',
                        '@cupcake', '@danish', '@donut', '@dragée', '@fruitcake', '@gingerbread', '@gummi', '@ice', '@jelly-o',
                        '@liquorice', '@macaroon', '@marzipan', '@oat', '@pie', '@plum', '@pudding', '@sesame', '@snaps', '@soufflé',
                        '@sugar', '@sweet', '@topping', '@wafer'
                    ],
                    minimumCharacters: 1
                }
            ]
        },
        removePlugins: [
            'ExportPdf',
            'ExportWord',
            'CKBox',
            'CKFinder',
            'EasyImage',
            'RealTimeCollaborativeComments',
            'RealTimeCollaborativeTrackChanges',
            'RealTimeCollaborativeRevisionHistory',
            'PresenceList',
            'Comments',
            'TrackChanges',
            'TrackChangesData',
            'RevisionHistory',
            'Pagination',
            'WProofreader',
            'MathType',
            'SlashCommand',
            'Template',
            'DocumentOutline',
            'FormatPainter',
            'TableOfContents',
            'PasteFromOfficeEnhanced'
        ]
    }).then(ckeditor => {
        console.log(Array.from(ckeditor.ui.componentFactory.names()));
        ckeditor1 = ckeditor
        ckeditor.plugins.get('FileRepository').createUploadAdapter = (loader) => {
            return {
                upload: () => {
                    return loader.file.then(file => {
                        return new Promise((resolve, reject) => {
                            const formData = new FormData();
                            formData.append('file', file);
                            formData.append('csrf', $("input[name='csrf']").val())
                            fetch('/channel/imageupload', {
                                method: 'POST',
                                body: formData
                            })
                                .then(response => response.json())
                                .then(data => {
                                    if (data.success) {
                                        resolve({
                                            default: data.url // URL to the uploaded image
                                        });
                                    } else {
                                        reject(data.error);
                                    }
                                })
                                .catch(error => {
                                    reject(error);
                                });
                        });
                    });
                }
            };
        };
    })
        .catch(error => {
            console.error(error);
        });

}


/** Page String */
function AddPageString(name, newid,languagedata) {

    var pagestr =
        `<div class="sort-index page0${newid}" data-id=0 data-newid=${newid} data-collapcount=${collapsecount}>
            <div class="spaceAccord accordionOption">
                <a class="spaceAccord-btn accordion-button viewpage" type="button" data-bs-toggle="collapse" data-bs-target="#accord${collapsecount}" aria-expanded="true" aria-controls="collapse${collapsecount}">
                    <div style="display:flex">
                    <p class="para">${name}</p>
                    <img src="/public/img/accordion-vector.svg" style="display:none;"/>
                    </div>
                    <div class="card-option field-delete">
                        <div class="btn-group language-group space-option-grp">
                            <button type="button" class="dropdown-tog" data-bs-toggle="dropdown" aria-expanded="false">
                                <img src="/public/img/card-option.svg" alt="">               
                            </button>
                            <ul class="dropdown-menu dropdown-menu-end" data-popper-placement="bottom-end" style=" position: absolute; inset: 0px 0px auto auto; margin: 0px; transform: translate(0px, 22px);">
                                <li>
                                    <button class="dropdown-item insertsubpage" type="button">
                                     ` + languagedata.Spaces.insertsubpages + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item pageabove" type="button">
                                        ` + languagedata.Spaces.addpageabove  +  `
                                    </button>
                                </li>
    
                                <li>
                                    <button class="dropdown-item pagebelow" type="button">
                                        `+ languagedata.Spaces.addpagebelow + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item rename" type="button">
                                        ` + languagedata.Spaces.rename + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item pageclone" type="button">
                                        ` + languagedata.Spaces.clone  + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item movepage" type="button" data-bs-toggle="modal" data-bs-target="#moveModal">
                                        ` + languagedata.Spaces.movepage + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item deletepage" type="button" data-bs-toggle="modal" data-bs-target="#centerModal">
                                        ` + languagedata.delete + `
                                    </button>
                                </li>
                            </ul>
                        </div>
                    </div>
                </a> 
            </div>
        </div>`

    collapsecount++;

    return pagestr
}

function AddGroupString(dbid,languagedata) {

    var dbid1 = dbid == "" ? 0 : dbid

    var grpstr = `
        <div class="groupsdiv" data-id=${dbid1} data-newid=${GroupId} style="margin-block:15px" id=grp${dbid1}${GroupId}>
            <div class="newAccord-group">
                <p class="para">GROUP${GroupId}</p>
                <div class="card-option field-delete">
                    <div class="btn-group language-group space-option-grp">
                        <button type="button" class="dropdown-tog" data-bs-toggle="dropdown" aria-expanded="false">
                            <img src="/public/img/card-option.svg" alt="">
                        </button>

                        <ul class="dropdown-menu dropdown-menu-end" data-popper-placement="bottom-end" style="position: absolute; inset: 0px 0px auto auto;margin: 0px;transform: translate(0px, 22px);">
                                <li>
                                    <button class="dropdown-item insertpage" type="button">
                                           `+languagedata.Spaces.insertpage + `
                                    </button>
                                </li>
                                <li>
                                    <button class="dropdown-item grouppageabove" type="button">
                                    ` + languagedata?.Spaces?.addpageabove  +  `
                                    </button>
                                </li>

                                <li>
                                    <button class="dropdown-item grouppagebelow" type="button">
                                    ` + languagedata.Spaces.addpagebelow + `
                                    </button>
                                </li>

                                <li>
                                    <button class="dropdown-item grouprename" type="button">
                                    ` + languagedata?.Spaces?.rename + `
                                    </button>
                                </li>
                                
                                <li>
                                    <button class="dropdown-item groupclone" type="button">
                                    ` + languagedata.Spaces.clone  + `
                                    </button>
                                </li>

                                <li>
                                    <button class="dropdown-item groupdelete" type="button">
                                    ` + languagedata.delete + `
                                    </button>
                                </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    `

    return grpstr
}

function AddSubString(colcount) {

    var substr =
        `
    <div id="accord${colcount}" class="accordion-collapse collapse show spaceAccord-content subpageview" aria-labelledby="headingOne" data-bs-parent="#accordionExample" data-id=0 data-newid=${SubPageId}>
        <p class="para high">NewPage${SubPageId}</p>
    </div>
    `
    return substr

}

/** View Page content */
$(document).on('click', '.viewpage', function () {

    var id = $(this).parents('.sort-index').attr('data-id');

    var newid = $(this).parents('.sort-index').attr('data-newid');

    $('.subpageview').removeClass('subpageselected');

    $('.viewpage').removeClass('pageselected');

    $(this).addClass('pageselected');

    const pageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })


    if (pageindex >= 0) {

        $('#pagetitle').val(newpagesarr[pageindex].Name)

        ckeditor1.setData(newpagesarr[pageindex].Content);

    }

})

/** View SubPage content */
$(document).on('click', '.subpageview', function () {

    var id = $(this).attr('data-id');

    var newid = $(this).attr('data-newid');

    $('.pageselected').removeClass('pageselected');

    $('.subpageview').removeClass('subpageselected');

    $(this).addClass('subpageselected');

    const pageindex = Subpagearr.findIndex(obj => {

        console.log(obj.SubPageId, id, "  ", obj.NewSubPageId, newid);

        return obj.SubPageId == id && obj.NewSubPageId == newid;

    })

    console.log(pageindex);

    if (pageindex >= 0) {

        console.log(Subpagearr[pageindex].Name);

        $('#pagetitle').val(Subpagearr[pageindex].Name)

        ckeditor1.setData(Subpagearr[pageindex].Content);

    }

})

/** Keyup content save to object */
$(document).on("DOMSubtreeModified", ".ck-restricted-editing_mode_standard", function () {

    var data = ckeditor1.getData();

    var pageid = $('.pageselected').parents('.sort-index').attr('data-id');

    var newpageid = $('.pageselected').parents('.sort-index').attr('data-newid');

    var subpageid = $('.subpageselected').attr('data-id');

    var newsubpageid = $('.subpageselected').attr('data-newid');

    if (pageid != 0 || newpageid != 0) {

        const pageindex = newpagesarr.findIndex(obj => {

            console.log(obj.PageId, pageid, "  ", obj.NewPageId, newpageid);

            return obj.PageId == pageid && obj.NewPageId == newpageid;

        })

        if (pageindex >= 0) {

            newpagesarr[pageindex].Content = data

        }

    }

    if (subpageid != 0 || newsubpageid != 0) {

        const pageindex = Subpagearr.findIndex(obj => {

            return obj.SubPageId == subpageid && obj.NewSubPageId == newsubpageid;

        })

        if (pageindex >= 0) {

            Subpagearr[pageindex].Content = data

        }

    }
})

/**Page above  */
$(document).on('click', '.pageabove', function () {

    var pagename = "New Page " + PageId

    var AddPage = AddPageString(pagename, PageId,languagedata);

    $(AddPage).insertBefore($(this).parents(".sort-index"));

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;
})

/**Page below  */
$(document).on('click', '.pagebelow', function () {

    var pagename = "New Page " + PageId

    var AddPage = AddPageString(pagename, PageId,languagedata);

    $(AddPage).insertAfter($(this).parents(".sort-index"));

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;
})

/** Move Page */
$(document).on('click','.movepage',function(){

    var grps=""

    for(let x of newGrouparr){

        grps +=`<button type="button"  class="dropdown-item">${x.Name}</button`
    }

    $('.movegroupdiv').html(grps);

})

/** Clone Page */
$(document).on('click', '.pageclone', function () {

    var id = $(this).parents('.sort-index').attr('data-id');

    var newid = $(this).parents('.sort-index').attr('data-newid');

    const pageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == id && obj.NewPageId == newid;

    })

    if (pageindex >= 0) {

        var AddPage = AddPageString(newpagesarr[pageindex].Name, PageId , languagedata);

        $(AddPage).insertAfter($(this).parents(".sort-index"));

        var pageobj = PageObjectCreate();

        pageobj.Name = newpagesarr[pageindex].Name

        pageobj.Content = newpagesarr[pageindex].Content

        newpagesarr.push(pageobj)

        const subfilter = Subpagearr.filter((obj) => obj.ParentId == id && obj.NewParentId == newid)

        for (let x of subfilter) {

            var colcount = $('.page0' +PageId).attr('data-collapcount')

            var Subpage = AddSubString(colcount);

            $('.page0' + PageId).children('.accordionOption').append(Subpage);

            $('.page0' +PageId).children('.spaceAccord').children('a').children('div').children('img').show();

            var obj = SubPageObjectCreate();

            obj.Name = "NewPage" + SubPageId
        
            Subpagearr.push(obj)
        
            SubPageId++;

        }

        PageId++;

        OrderIndex++;
    }

})

/** Insert Page into group */
$(document).on('click', '.insertpage', function () {

    var id = $(this).parents('.groupsdiv').attr('data-id');

    var newid = $(this).parents('.groupsdiv').attr('data-newid');

    var pagename = "New Page " + PageId

    var pagedata = AddPageString(pagename, PageId , languagedata)

    $('#grp' + id + newid).append(pagedata)

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;

})

/** keyup name change */
$(document).on('keyup', "#pagetitle", function () {

    var id = $('.pageselected').parents('.viewpage').attr('data-id');

    var newid = $('.pageselected').parents('.viewpage').attr('data-newid');

    var subid = $('.subpageselected').attr('data-id');

    var newsubid = $('.subpageselected').attr('data-newid')

    var value = $(this).val();

    if (id != 0 || newid != 0) {

        $('.pageselected').children('div').children('p').text(value);

        const pageindex = newpagesarr.findIndex(obj => {

            return obj.PageId == id && obj.NewPageId == newid;

        })

        if (pageindex >= 0) {

            newpagesarr[pageindex].Name = value

        }
    }

    if (subid != 0 || newsubid != 0) {

        $('.subpageselected').children('p').text(value);

        const pageindex = Subpagearr.findIndex(obj => {

            return obj.SubPageId == subid && obj.NewSubPageId == newsubid;

        })

        if (pageindex >= 0) {

            Subpagearr[pageindex].Name = value

        }

    }

})

/*Delete page Confirmation */
$(document).on('click', '.deletepage', function () {

    var pageid = $(this).parents(".sort-index").attr("data-id")

    var newpgid = $(this).parents(".sort-index").attr("data-newid")

    const pageindex = newpagesarr.findIndex(obj => {

        return obj.PageId == pageid && obj.NewPageId == newpgid;

    })

    if (pageindex >= 0) {

        $("#delid").attr("data-pageid", pageid)

        $('#delid').attr('data-newpageid', newpgid)

        $('.deltitle').text(languagedata.Spaces.delpagetitle)

        $('.deldesc').text(languagedata.Spaces.delpagecontent)

        $('.delname').text(newpagesarr[pageindex].Name)

        $('#content').text(languagedata.Spaces.delpagecontent)

    }

})

/** Delete Page & Group & Subpage */
$(document).on('click', '#delid', function () {

    var pageid = $(this).attr('data-pageid');

    var newpageid = $(this).attr('data-newpageid');

    var groupid = $(this).attr('data-groupid');

    var newgroupid = $(this).attr('data-newgroupid');

    var subpageid = $(this).attr('data-subpageid');

    var newsubpageid = $(this).attr('data-newsubpageid');

    if (pageid != undefined || newpageid != undefined) {

        for (let x of Subpagearr) {

            if (x.ParentId) {

            }

        }

        const pageindex = newpagesarr.findIndex(obj => {

            console.log(obj.PageId, pageid, obj.NewPageId, newpageid);

            return obj.PageId == pageid && obj.NewPageId == newpageid;

        })

        if (pageindex >= 0) {

            console.log("asfd");

            newpagesarr.splice(pageindex, 1)

            $('.page' + pageid + newpageid).remove();

        }

    }

    if (groupid != undefined || newgroupid != undefined) {

        const groupindex = newGrouparr.findIndex(obj => {

            return obj.GroupId == groupid && obj.NewGroupId == newgroupid;

        })

        if (groupindex >= 0) {

            newGrouparr.splice(groupindex, 1)

        }

    }

    $('#delcancel').trigger('click');
})


/**Add Page */
$(document).on('click', '#addpage', function () {

    var pagename = "New Page " + PageId

    var pagedata = AddPageString(pagename, PageId,languagedata)

    $('.newpages').append(pagedata)

    var pageobj = PageObjectCreate();

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;

})

$(document).on('click', '#addgroup', function () {

    var grpstring = AddGroupString(0,languagedata)

    $('.newpages').append(grpstring)

    var obj = GroupCreationObject()

    obj.Name = "Group" + GroupId

    newGrouparr.push(obj)

    var pagename = "New Page " + PageId

    var pagedata = AddPageString(pagename, PageId ,languagedata)

    $('#grp' + 0 + GroupId.toString()).append(pagedata)

    var pageobj = PageObjectCreate()

    pageobj.Name = pagename

    newpagesarr.push(pageobj)

    PageId++;

    OrderIndex++;

    GroupId++;

})

/** Page Creation Object */
function PageObjectCreate() {

    var obj = {}

    obj.PageId = 0,

        obj.NewPageId = PageId,

        obj.Content = "",

        obj.PageGroupId = 0,

        // obj.NewGroupId=parseInt(newgrpid==undefined?0:newgrpid),

        obj.OrderIndex = OrderIndex,

        obj.ParentId = 0

    return obj

}


/** SubPage Creation Object */
function SubPageObjectCreate() {

    var obj = {}

        obj.SubPageId = 0,

        obj.NewSubPageId = SubPageId,

        obj.Content = "",

        obj.PageGroupId = 0,

        // obj.NewGroupId = parseInt(newgrpid == undefined ? 0 : newgrpid),

        obj.OrderIndex = OrderIndex

    // obj.ParentId= parseInt(parentid)

    return obj

}


/** Group Creation Object */
function GroupCreationObject() {

    var obj = {}

    obj.GroupId = 0,

        obj.NewGroupId = GroupId,

        //  obj.Name = name.toUpperCase(),

        obj.OrderIndex = OrderIndex

    return obj

}



// Savebtn
$(document).on('click', '#savebtn', function () {

    var spid = $("#spid").val();

    $.ajax({
        url: "/spaces/page/pagecreate",
        type: "POST",
        dataType: "json",
        data: {
            "spaceid": spid,
            "creategroups": JSON.stringify({ newGrouparr }),
            "createpages": JSON.stringify({ newpagesarr }),
            "createsubpage": JSON.stringify({ Subpagearr }),
            "deletepage": JSON.stringify({ deletePagearr }),
            "deletegroup": JSON.stringify({ deletegrouparr }),
            "deletesub": JSON.stringify({ deletesubarr }),
            "save": "save",
            csrf: $("input[name='csrf']").val()
        },
        success: function (result) {
            if (result) {

                setCookie('get-toast', languagedata.Toast.pagesupdatedsuccessfully, 1)
                setCookie('Alert-msg', 'success', 1)

            } else {

                setCookie('Alert-msg', languagedata.Toast.internalserverr, 1)
            }

            // window.location.reload(true);

        }
    })
})

/*Insert Sub Page */
$(document).on('click', '.insertsubpage', function () {

    var colcount = $(this).parents('.sort-index').attr('data-collapcount')

    var Subpage = AddSubString(colcount);

    $(this).parents('.spaceAccord').append(Subpage);

    var obj = SubPageObjectCreate();

    obj.Name = "NewPage" + SubPageId

    Subpagearr.push(obj)

    console.log(Subpagearr);

    SubPageId++;

    obj.ParentId = parseInt($(this).parents('.sort-index').attr('data-id'));

    obj.NewParentId = parseInt($(this).parents('.sort-index').attr('data-newid'));

    $(this).parents('.spaceAccord-btn').children('div').children('p').siblings('img').show();

})

/** revision close */
$(document).on('click','#revisionclose',function(){

    $(this).parents('.space-section-right').hide();
})

/**revision open */
$(document).on('click','.revision-btn',function(){
    
    $('.space-section-right').show();

})