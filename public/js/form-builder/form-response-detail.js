$(document).ready(function () {


    const quill = new Quill('#editor', {
        theme: 'snow',
        modules: {
            toolbar: [
                ['bold', 'italic', 'underline'],
                [{ list: 'ordered' }, { list: 'bullet' }]
            ]
        }
    });


    $(".ql-snow").addClass("hidden")

    $("#replyButton").click(function () {
        $(".ql-snow").removeClass("hidden")
        $("#editor").removeClass("hidden")
        $("#editorSaveandCancel").removeClass("hidden")
        $("#editorSubmitBtn").removeClass("hidden")
        $("#quickMessages").removeClass("hidden")
    })


    $("#editorSubmitBtn").click(function () {

        const text = quill.getText().trim();

        if (!text) {

            notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  flex relative max-sm:max-w-[300px] items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-error.svg" alt = "toast error"></div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >Please Enter the Content</p ></div ></div ></li></ul> `;
            $(notify_content).insertBefore(".editorerror");
            setTimeout(function () {
                $('.toast-msg').fadeOut('slow', function () {
                    $(this).remove();
                });
            }, 5000); // 5000 milliseconds = 5 seconds
            return;
        }

        let email = $("#editorEmail").val()

        let ticket = $("#ticketNumber").val()

        let username = $("#UserName").val()

        const htmlContent = quill.root.innerHTML;


        $.ajax({
            url: "/admin/cta/replyforresponse",
            type: "POST",
            async: false,
            data: { csrf: $("input[name='csrf']").val(), "htmlContent": htmlContent, "email": email, "ticket": ticket, "username": username },
            datatype: "json",
            caches: false,
            success: function (data) {


                if (data.status == true) {

                    setCookie("get-toast", "Reply Submitted Successfully")

                    window.location.href = "/admin/cta/form-responses"

                }
            }
        })


    })

    $(document).on("click", ".quickMessages", function () {

        let value = $(this).text().trim();

        let length = quill.getLength();
        if (length > 1) {
            quill.insertText(length - 1, "\n");
        }

        let templatevalue = "";

        if (value === "General Acknowledgement") {
            templatevalue = [
                "Hello,",
                "",
                "Thank you for contacting Piccosoft.",
                "We’ve received your message and our team is reviewing it. We’ll get back to you shortly with an update.",
                "If you need immediate assistance, feel free to reply to this email.",
                "",
                "Best regards,",
                "Piccosoft Support Team"
            ].join("\n");

        } else if (value === "Product / Service Enquiry Reply") {
            templatevalue = [
                "Hello,",
                "",
                "Thank you for reaching out to Piccosoft.",
                "We’ve received your enquiry regarding our products/services.",
                "Our team will review the details and get back to you shortly.",
                "",
                "Best regards,",
                "Piccosoft Team"
            ].join("\n");

        } else if (value === "Support Request Acknowledgement") {
            templatevalue = [
                "Hello,",
                "",
                "Thanks for contacting Piccosoft Support.",
                "Your request has been received and forwarded to our support team.",
                "We’ll update you as soon as possible.",
                "",
                "Best regards,",
                "Piccosoft Team"
            ].join("\n");

        } else if (value === "Sales / Demo / Pricing Enquiry") {
            templatevalue = [
                "Hello,",
                "",
                "Thank you for your interest in Piccosoft.",
                "Our sales team will reach out shortly with pricing or demo details.",
                "",
                "Best regards,",
                "Piccosoft Sales Team"
            ].join("\n");

        } else if (value === "Request for More Information") {
            templatevalue = [
                "Hello,",
                "",
                "Thank you for contacting Piccosoft.",
                "Could you please share a few more details regarding your request?",
                "Once we have the information, we’ll proceed accordingly.",
                "",
                "Best regards,",
                "Piccosoft Team"
            ].join("\n");

        } else if (value === "Closure / Resolution Message") {
            templatevalue = [
                "Hello,",
                "",
                "We’re happy to inform you that your request has been resolved.",
                "If you need any further assistance, feel free to reach out.",
                "",
                "Warm regards,",
                "Piccosoft Team"
            ].join("\n");
        }

        quill.insertText(quill.getLength() - 1, templatevalue);
        quill.setSelection(quill.getLength(), 0);
        quill.focus();
    });

})