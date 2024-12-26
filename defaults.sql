-- Must use separate insert queries for parent and child data insertions to maintain sequence.

-- Some reserved words are used for dynamic values insertion in queries.They are given as below.Use them whenever they needed essentially.

-- current-time,uid(user id),tid(tenant id),pid(parent id),chid(channel id),blid(block id),tagid(tag id),mapcat(mapped categories)

INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Default Category', 'default_category', 'current-time', uid, 0, 0, 'Default_Category', tid);

INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Default1', 'default1', 'current-time', uid, 0, pid, 'Default1',tid);

INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('Default_Channel', 'default_channel', 0, 1, 0, 'current-time', uid, 'default description','mychannels',tid);

INSERT INTO tbl_channel_categories(channel_id, category_id, created_at, created_on,tenant_id) VALUES (chid,'mapcat', uid, 'current-time',tid);

INSERT INTO tbl_member_groups(name,slug,description,is_active,is_deleted,created_on,created_by,tenant_id) VALUES ('Default Group', 'default-group', 'Default Group', 1,0, 'current-time', uid,tid);

INSERT INTO tbl_member_settings (allow_registration,member_login,notification_users,tenant_id) VALUES (1,'password','uid',tid)

INSERT INTO tbl_email_configurations (selected_type,tenant_id) VALUES ('environment',tid)

INSERT INTO tbl_general_settings (date_format,expand_logo_path,language_id,logo_path,storage_type,tenant_id,time_format,time_zone) VALUES ('dd mmm yyyy','/public/img/logo-bg.svg',1,'/public/img/logo1.svg','aws',tid,'12','Asia/Kolkata')

INSERT INTO tbl_graphql_settings (token_name,description,duration,created_by,created_on,is_deleted,token,is_default,tenant_id) VALUES ('Default Token','Default Token','Unlimited',uid,'current-time',0,'apikey',1,tid)

-- below insert queries are for CREATING DEFAULT CATEGORIES, CHANNELS and TEMPLATES 

-- ****************** Knowledge-Base ***********************
--  default knowledge base categories for channel based templates
INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by, is_deleted, parent_id, description, tenant_id) VALUES ('Knowledge Base', 'knowledge-base', 'current-time', uid, 0,0,'Parent Category for knowledge base', tid)

-- default knowledge base channels for channel based templates
INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('Knowledge Base', 'knowledge-base', 0, 1, 0, 'current-time', uid, 'Channel for knowledge base templates','mychannels',tid);

INSERT INTO tbl_module_permissions(route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES ('rname','Knowledge Base','knowledge-base',mid,1,0,1,2,uid,'current-time',tid)

-- INSERT INTO tbl_channel_categories(channel_id, category_id, created_at, created_on,tenant_id) VALUES (chid,'mapcat', uid, 'current-time',tid);

-- default knowledge-base template module creation
INSERT INTO tbl_template_modules(template_module_name,template_module_slug,description,is_active,created_by,created_on,channel_id,tenant_id) VALUES ( 'Help Center', 'help-center', 'Help Center Templates', 1, uid, 'current-time', chid,tid)

-- default knowledge-base template creation
INSERT INTO tbl_templates(template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES('Knowledge Base Theme','knowledge-base-theme', tempmoduleid,'knowledgeBase.png','media/knowledgeBase.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/knowledge-base&demo-title=knowledge-base&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/knowledge-base','current-time',uid,1,0,chid,tid)


-- ***************************** Blogs *********************************
-- Please DON'T change the flow of the input statements it will definitely affect the output.

--  default Parent blog CATEGORY for channel based templates
INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by, is_deleted, parent_id, description, tenant_id) VALUES ('Blog', 'blog', 'current-time', uid, 0,0,'Parent Category for blogs', tid)

-- default Parent blog CHANNEL for channel based templates
INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('Blog', 'blog', 0, 1, 0, 'current-time', uid, 'Channel for blog templates','mychannels',tid);

INSERT INTO tbl_module_permissions(route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES ('rname','Blog','blog',mid,1,0,1,2,uid,'current-time',tid)


-- default blog template module creation
INSERT INTO tbl_template_modules(template_module_name,template_module_slug,description,is_active,created_by,created_on,channel_id,tenant_id) VALUES ( 'Blog', 'blog', 'Blog Templates', 1, uid, 'current-time',chid,tid )

-- default NextJS Starter Theme template creation
INSERT INTO tbl_templates( template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES('NextJS Starter Theme','next-js-starter-theme', tempmoduleid,'blogstarter.png','media/blogstarter.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-starter-theme&demo-title=nextjs-starter-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-starter-theme','current-time',uid,1,0,chid,tid)

-- default NextJS Blog Theme template creation
INSERT INTO tbl_templates(template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES('NextJS Blog Theme','next-js-blog-theme', tempmoduleid,'blog1.png','media/blog1.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog-theme&demo-title=nextjs-blog-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog-theme','current-time',uid,1,0,chid,tid)


-- default NextJS Blogs2 Theme template creation
INSERT INTO tbl_templates(template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES('NextJS Blog2 Theme','next-js-blog2-theme', tempmoduleid,'blog2.png','media/blog2.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog2-theme&demo-title=nextjs-blog2-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog2-theme','current-time',uid,1,0,chid,tid)


-- default NextJS Blogs4 Theme template creation
INSERT INTO tbl_templates( template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,channel_id,tenant_id) VALUES('NextJS Blog4 Theme','next-js-blog4-theme', tempmoduleid,'blog4.png','media/blog4.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog4-theme&demo-title=nextjs-blog4-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog4-theme','current-time',uid,1,0,chid,tid)




