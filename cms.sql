--Insert Default Values

INSERT INTO tbl_roles(id, name, description, slug, is_active, is_deleted, created_on, created_by) VALUES (1, 'Super Admin', 'Default user type. Has the full administration power', 'super_admin', 1, 0, '2023-07-25 05:50:14', 1);

INSERT INTO tbl_roles(id, name, description, slug, is_active, is_deleted, created_on, created_by,tenant_id) VALUES (2, 'Admin', 'Default user type. Has the full administration power', 'admin', 1, 0, '2025-03-14 11:09:12', 1,1);

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (1,'English', 'en', '/public/img/in.jpeg','locales/en.json', 1, 1,1, 'current_time',0)

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (2,'Spanish', 'es','/public/img/sp.jpeg', 'locales/es.json', 1, 0,1, 'current_time',0)

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (3,'French', 'fr','/public/img/fr.jpeg', 'locales/fr.json', 1, 0,1, 'current_time',0)

INSERT INTO tbl_languages(id,language_name,language_code,image_path,json_path,is_status,is_default,created_by,created_on,is_deleted) VALUES (4,'Russian', 'ru','/public/img/russia.svg', 'locales/ru.json', 1, 0,1, 'current_time',0)

INSERT INTO tbl_users(id, role_id, first_name, last_name, email, username, password, mobile_no, is_active, created_on, created_by, is_deleted, default_language_id,tenant_id)	VALUES (1, 2, 'Spurt', 'Cms', 'spurtcms@gmail.com', 'Admin', '$2a$14$3kdNuP2Fo/SBopGK9e/9ReBa8uq82vMM5Ko2jnbLpuPb6qxqgR0x2', '9999999999', 1, '2023-11-24 14:56:12',1, 0, 1,1);

INSERT INTO tbl_user_personalizes(id,user_id,menu_background_color,logo_path,expand_logo_path) values(1,1,'rgb(9, 171, 217)','/public/img/logo.svg','/public/img/logo-bg.svg');

INSERT INTO tbl_general_settings(id, logo_path, expand_logo_path, date_format, time_format,language_id,time_zone,tenant_id) VALUES (1,'/public/img/logo1.svg', '/public/img/logo-bg.svg', 'dd mmm yyyy', '12',1,'Asia/Kolkata',1);

INSERT INTO tbl_email_configurations (id,selected_type) VALUES (1,'environment')

INSERT INTO tbl_storage_types(id, local, selected_type) VALUES (1, 'storage', 'aws');

--Channel default fields

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (1, 'Label', 'label', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (2, 'Text', 'text', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (3, 'Link', 'link', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (4, 'Date & Time', 'date&time', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (5, 'Select', 'select', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (6, 'Date', 'date', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (7, 'TextBox', 'textbox', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (8, 'TextArea', 'textarea', 1, 0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (9, 'Radio Button', 'radiobutton', 1, 0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (10, 'CheckBox', 'checkbox', 1, 0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (11, 'Text Editor', 'texteditor', 1, 0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (12, 'Section', 'section', 1, 0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (13, 'Section Break', 'sectionbreak', 1, 0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (14, 'Members', 'member', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (15, 'Media Gallery', 'mediagallrey', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (16, 'Video URL', 'videourl', 1,  0, 1, '2023-03-14 11:09:12');

INSERT INTO tbl_field_types(id, type_name, type_slug, is_active, is_deleted, created_by, created_on) VALUES (17, 'Number', 'number', 1,  0, 1, '2023-03-14 11:09:12');

--Default Insert Menu value

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES (1, 'Dashboard', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/dashboard_new.svg', '', 1, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES (2, 'Categories', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/category_new.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 2, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(3, 'Channels', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/cms-on.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 3, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(4, 'Entries', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/left-entries.svg', 'Seamlessly add and update content within specific channels for targeted delivery and right audience. ', 4, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(5, 'Media', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/media_1.svg', 'Manage and integrate media files effortlessly to enrich your content.', 5, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(6, 'Users', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/member_new.svg', 'Admin manages user profiles, including adding new users and assigning them to specific groups.', 8, 'left',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(7, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 0, 1, '/public/img/settings.svg', 'Check relevant boxes for providing access control to this role, in this settings.', 9, 'left',0,1);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(8, 'Channels', 1, 1, '2023-03-14 11:09:12', 0, 3, 0, '/public/img/accord-channels.svg', 'Easily create and manage diverse content types using Spurt CMS Admin.', 10, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(9, 'Entries', 1, 1, '2023-03-14 11:09:12', 0, 4, 0, '/public/img/entries.svg', 'Seamlessly add and update content within specific channels for targeted delivery and right audience. ', 11, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(10, 'Categories Group', 1, 1, '2023-03-14 11:09:12', 0, 2, 0, '/public/img/categories-group.svg', 'Create various categories for content and organize them efficiently by categorizing it with user-defined labels.', 12, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(11, 'Categories', 1, 1, '2023-03-14 11:09:12', 0, 2, 0, '/public/img/categories-group.svg', '', 13, 'tab',0,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(12, 'Users', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/accord-member.svg', 'Add member profiles, map them to a member group and manage the entire list of member profiles. ', 14, 'tab',0,0);

Update tbl_modules Set module_name='Users' where id=12

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(13, 'Users Group', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/group.svg', 'Create groups and categorize members into various groups like Elite Members or Favorite Members', 15, 'tab',0,0);

Update tbl_modules Set module_name='Users Group' where id=13

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(14, 'Member Restrict', 1, 1, '2023-03-14 11:09:12', 0, 6, 0, '/public/img/accord-memberRestrict.svg', '', 16, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(15, 'My Account', 1, 1, '2023-03-14 11:09:12', 0, 7, 1, '/public/img/profile.svg', 'Here you can change your email id and password.', 17, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(16, 'Personalize', 1, 1, '2023-03-14 11:09:12', 0, 7, 1, '/public/img/personalize.svg', 'You can change the theme according your preference. ', 18, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(17, 'Security', 1, 1, '2023-03-14 11:09:12', 0, 7, 1, '/public/img/security.svg', 'To protech your account, you can change your password at regular intervals.', 19, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(18, 'Roles & Permissions', 1, 1, '2023-03-14 11:09:12', 0, 7, 0, '/public/img/roles-permision.svg', 'Create different admin roles and assign permissions to them based on their responsibilities.', 20, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(19, 'Team', 1, 1, '2023-03-14 11:09:12', 0, 7, 0, '/public/img/team.svg', 'Create a team of different admin roles.', 21, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(21, 'Email', 1, 1, '2023-03-14 11:09:12', 0, 7, 0, '/public/img/email.svg', 'Create and Maintain Email Templates that are to be sent to users in different scenarios.', 22, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(22, 'Data', 1, 1, '2023-03-14 11:09:12', 0, 7, 0, '/public/img/settings.svg', '', 23, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(23, 'Dashboard', 1, 1, '2023-03-14 11:09:12', 0, 1, 1, '/public/img/settings.svg', '', 24, 'tab',1,0);

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(24, 'Media', 1, 1, '2023-03-14 11:09:12', 0, 5, 1, '/public/img/media.svg', '', 25, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(25, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 7, 1, '/public/img/settings.svg', '', 26, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(26, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 3, 1, '/public/img/settings.svg', '', 27, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(27, 'Settings', 1, 1, '2023-03-14 11:09:12', 0, 5, 1, '/public/img/settings.svg', '', 28, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(28, 'Member Settings', 1, 1, '2023-03-14 11:09:12', 0, 6, 1, '/public/img/settings.svg', '', 29, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(29, 'General Settings', 1, 1, '2023-03-14 11:09:12', 0, 7, 0, '/public/img/settings.svg', '', 30, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(30, 'Graphql Playground', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/graphql-playground.svg', '', 31, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id,assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(31, 'Graphql Playground', 1, 1, '2023-03-14 11:09:12', 0, 30, 0, '/public/img/graphql-playground.svg', '', 32, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission,icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(32, 'Document', 1, 1, '2023-03-14 11:09:12', 0, 30, 0, '/public/img/settings.svg', '', 33, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(33, 'Blocks', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/Blocks.svg', '', 34, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(34, 'Blocks', 1, 1, '2023-03-14 11:09:12', 0, 33, 0, '/public/img/Blocks.svg', '', 35, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(35, 'Next.js Templates', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/templates.svg', '', 36, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(36,  'Next.js Templates', 1, 1, '2023-03-14 11:09:12', 0, 35, 0, '/public/img/templates.svg', '', 37, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(37, 'Graphql Api', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/role-graph-ql.svg', '', 38, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(38, 'Graphql Api', 1, 1, '2023-03-14 11:09:12', 0, 37, 0, '/public/img/Graph-QL.svg', '', 39, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(39, 'Webhooks', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/Webhooks.svg', '', 40, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(40, 'Webhooks', 1, 1, '2023-03-14 11:09:12', 0, 39, 0, '/public/img/Webhooks.svg', '', 41, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(41, 'Languages', 1, 1, '2023-03-14 11:09:12', 0, 7, 0, '/public/img/language.svg', '', 42, 'tab',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(42, 'Membership', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/Membership.svg', '', 6, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(43, 'Membership', 1, 1, '2023-03-14 11:09:12', 0, 42, 0, '/public/img/Membership.svg', '', 7, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(44, 'Forms Builder', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/Group 17556.svg', '', 43, 'left',1,1)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(45, 'Forms Builder', 1, 1, '2023-03-14 11:09:12', 0, 44, 0, '/public/img/Group 17556.svg', '', 44, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(46, 'AI Settings', 1, 1, '2023-03-14 11:09:12', 0, 7, 0, '/public/img/open-ai.svg', '', 45, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(47, 'Courses', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/Course black.svg', '', 46, 'left',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(48, 'Courses', 1, 1, '2023-03-14 11:09:12', 0, 47, 0, '/public/img/Course black.svg', '', 47, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(49, 'Integration', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/integration.svg', '', 48, 'left',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(50, 'Integration', 1, 1, '2023-03-14 11:09:12', 0, 49, 0, '/public/img/integration.svg', '', 49, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(51, 'Listing', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/Default Listing.svg', '', 50, 'left',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(52, 'Listing', 1, 1, '2023-03-14 11:09:12', 0, 51, 0, '/public/img/Default Listing.svg', '', 51, 'tab',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(53, 'Website', 1, 1, '2023-03-14 11:09:12', 0, 0, 0, '/public/img/Default state.svg', '', 52, 'left',1,0)

INSERT INTO tbl_modules(id, module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) VALUES(54, 'Website', 1, 1, '2023-03-14 11:09:12', 0, 53, 0, '/public/img/Default state.svg', '', 53, 'tab',1,0)

UPDATE tbl_modules set module_name='Website' WHERE  id IN (53,54)
--Default Module Permission Routes

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (1, '/dashboard', 'Dashboard', '', 23, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'dashboard')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (2, '/categories/', 'View', 'Give view access to the category group', 10, 1, '2023-03-14 11:09:12', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (3, '/categories/newcategory', 'Create', 'Give create access to the category Group', 10, 1, '2023-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (4, '/categories/updatecategory', 'Update', 'Give update access to the category Group', 10, 1, '2023-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (5, '/categories/deletecategory', 'Delete', 'Give delete access to the category Group', 10, 1, '2023-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (6, '/categories/addcategory/','View','Give view access to the categories', 11, 1, '2023-03-14 11:09:12', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (7, '/categories/addsubcategory', 'Create', 'Give create access to the categories', 11, 1, '2023-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (8, '/categories/editsubcategory', 'Update', 'Give update access to the categories', 11, 1, '2023-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (9, '/categories/removecategory', 'Delete', 'Give delete access to the categories', 11, 1, '2023-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (10, '/channels/', 'Channel', 'Give full access to the channels', 8, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'channels')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (11, '/usergroup/', 'View', 'Give view access to the member group', 13, 1, '2023-03-14 11:09:12', 0, 0,1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (12, '/usergroup/newgroup', 'Create', 'Give create access to the member group', 13, 1, '2023-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (13, '/usergroup/updategroup', 'Update', 'Give update access to the member group', 13, 1, '2023-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (14, '/usergroup/deletegroup', 'Delete', 'Give delete access to the member group', 13, 1, '2023-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (15, '/user/', 'View', 'Give view access to the Member', 12, 1, '2023-03-14 11:09:12', 0, 0, 1, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (16, '/user/newmember', 'Create', 'Give create access to the Member', 12, 1, '2023-03-14 11:09:12', 0, 0, 1, 2, 'create')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (17, '/user/updatemember', 'Update', 'Give update access to the Member', 12, 1, '2023-03-14 11:09:12', 0, 0, 1, 3, 'update')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (18, '/user/deletemember', 'Delete', 'Give delete access to the Member', 12, 1, '2023-03-14 11:09:12', 0, 0, 1, 4, 'delete')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (19, '/memberaccess/', 'Member-Restrict', 'Create and manage content consumption access and restrictions to the Members on the website. ', 14, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'memberaccess')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (20, '/entries/entrylist', 'Entries', '', 9, 1, '2023-03-14 11:09:12', 1, 0, 0, 0, 'entries')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (21, '/media/', 'View', '', 24, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'view')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (22, '/settings/roles/', 'Roles & Permission', 'Create and manage additional roles based on the needs for their tasks and activities in the control panel.', 18, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'roles')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (23, '/settings/users/userlist', 'Teams', 'Add users for assigning them roles, followed by assigning permissions for appropriate access to the CMS.', 19, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'users')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (25, '/settings/emails', 'Email Templates', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 21, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'email')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (26, '/settings/data', 'Data', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 22, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'data')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (27, '/categories/settings', 'Settings', '', 0, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (28, '/entries/settings', 'Settings', '', 26, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (29, '/media/settings', 'Settings', '', 27, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (30, '/user/settings', 'Settings', '', 28, 1, '2023-03-14 11:09:12', 1, 0, 0, 1, 'settings')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (31, '/settings/general-settings/', 'General Settings', 'Add different languages here, for the Members to consume content in the language of preference on the website.', 29, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'data')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by,created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (33, '/graphql', 'Graphql Api', 'Give full access to the graphql api', 38, 1, '2023-03-14 11:09:12', 1, 0, 1, 1, 'graphql')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES(35, '/templates', 'Template', 'Give full access to the template', 36, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'template')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (36, '/blocks', 'Blocks', 'Give full access to the blocks', 34, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'blocks')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (37, '/graphql/playground', 'Graphql Playground', 'Give full access to the graphql playground', 31, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'graphqlplayground')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (38, '/webhooks', 'Webhooks', 'Give full access to the webhooks', 40, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'webhooks')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (39, '/settings/languages', 'Languages', 'Give full access to the languages', 41, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'languages')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (40, '/members', 'Membership', 'Give full access to the membership', 43, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'membership')

INSERT INTO tbl_module_permissions(id, route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (41, '/formsbuilder', 'Forms Builder', 'Give full access to the form builder', 45, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'forms-builder')

INSERT INTO tbl_module_permissions(id,route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (42,'/settings/aisettings/', 'AI Settings', 'Give full access to the AI Settings', 46, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'aimodulesettings')

INSERT INTO tbl_module_permissions(id,route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (43,'/courses', 'Courses', 'Give full access to the Courses', 48, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'courses')

INSERT INTO tbl_module_permissions(id,route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (44,'/integration', 'Integration', 'Give full access to the Integration', 50, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'integration')

DELETE FROM tbl_module_permissions WHERE id = 45;
INSERT INTO tbl_module_permissions(id,route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (45,'/listing', 'Listing', 'Give full access to the Listing', 52, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'listing')

INSERT INTO tbl_module_permissions(id,route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name) VALUES (46,'/website', 'Website', 'Give full access to the Content Website', 54, 1, '2023-03-14 11:09:12', 1, 0, 1, 2, 'website')


--default ai prompts


INSERT INTO tbl_apps(id, title, description,field_json_path,icon_path,icon_name, created_by, created_on, is_active,channel_slug) VALUES (1, 'Long blog article', 'Generate a blog article from topic or keywords based on word count and tone','aiassistance/blog.json', '/public/img/text-img.svg', 'text-img.svg', 1, '2023-03-14 11:09:12',1,'blog')

INSERT INTO tbl_apps(id, title, description,field_json_path,icon_path,icon_name, created_by, created_on, is_active,channel_slug) VALUES (2, 'Knowledge base', 'Knowledge bases reduce the time employees spend searching for information, allowing them to focus on their core tasks.','aiassistance/blog.json', '/public/img/text-img.svg', 'text-img.svg', 1, '2023-03-14 11:09:12',1,'knowledge-base')

INSERT INTO tbl_apps(id, title, description,field_json_path,icon_path,icon_name, created_by, created_on, is_active,channel_slug) VALUES (3, 'News', 'In an age of information overload, news organizations are prioritizing in-depth reporting and investigative journalism to stand out','aiassistance/blog.json', '/public/img/text-img.svg', 'text-img.svg', 1, '2023-03-14 11:09:12',1,'news')

INSERT INTO tbl_apps(id, title, description,field_json_path,icon_path,icon_name, created_by, created_on, is_active,channel_slug) VALUES (4, 'Ecommerce', 'Innovations in technology, such as faster internet speeds, mobile shopping apps, and secure payment systems, have made online shopping more accessible and convenient.','aiassistance/blog.json', '/public/img/text-img.svg', 'text-img.svg', 1, '2023-03-14 11:09:12',1,'ecommerce')

INSERT INTO tbl_apps(id, title, description,field_json_path,icon_path,icon_name, created_by, created_on, is_active,channel_slug) VALUES (5, 'Jobs', 'Proficiency in technology-related fields, such as data analysis, coding, and cybersecurity, is in high demand. Workers who can navigate digital tools and platforms will have a competitive edge','aiassistance/blog.json', '/public/img/text-img.svg', 'text-img.svg', 1, '2023-03-14 11:09:12',1,'jobs')

INSERT INTO tbl_ai_prompts(id, app_id, master_id, child_id,prompt_level,system_prompt,user_prompt, created_by, created_on) VALUES (1, 1, 1,null,1, 'You are the best content creator for a common blog website', 'Generate a blog article for the given topic: {topic} in {articleLength}. The article should be in the tone of {tone}. The article should be based on the keywords: "{keywords}", and you may consider those keywords as sub-concepts. Generate the article effectively and exactly in the {tone} tone. The output should be in the format of innerHTML with Tailwind CSS and the response doesn''t contain any additional texts or notes', 1, '2023-03-14 11:09:12')

INSERT INTO tbl_ai_prompts(id, app_id, master_id, child_id,prompt_level,system_prompt,user_prompt, created_by, created_on) VALUES (2, 1, 1,1,1, 'You are the best content creator for a common blog website', 'with the given topic: {topic}. Generate the following contents as the brief details for the blog article. title - title should be as same as the topic, keywords - create 5 keyword that helps to filter the blog with the user keywords. articleLength - the maximum content length in words. reply only with json: { topic: "topic",keywords:["keywords"],articleLength:"count of words"}', 1, '2023-03-14 11:09:12')

INSERT INTO tbl_ai_prompts(id, app_id, master_id, child_id,prompt_level,system_prompt,user_prompt, created_by, created_on) VALUES (3, 1, 1,2,1, 'You are the best content creator for a common blog website', 'I have the certain keyword to generate a topic for a blog article. Create 5 effective article topics based on the given keyword: {keyword} . the topics should be more robust and it should be in maximum 20 characters. reply in json format with the following structure: { topics: ["topics"]}', 1, '2023-03-14 11:09:12')


INSERT INTO tbl_timezones(id,timezone) VALUES (1,'Africa/Cairo'),(2,'Africa/Johannesburg'),(3,'Africa/Lagos'),(4,'Africa/Nairobi'),(5,'America/Argentina/Buenos_Aires'),(6,'America/Chicago'),(7,'America/Denver'),(8,'America/Los_Angeles'),(9,'America/Mexico_City'),(10,'America/New_York'),(11,'America/Sao_Paulo'),(12,'Asia/Bangkok'),(13,'Asia/Dhaka'),(14,'Asia/Dubai'),(15,'Asia/Hong_Kong'),(16,'Asia/Jakarta'),(17,'Asia/Kolkata'),(18,'Asia/Manila'),(19,'Asia/Seoul'),(20,'Asia/Shanghai'),(21,'Asia/Singapore'),(22,'Asia/Tokyo'),(23,'Australia/Melbourne'),(24,'Australia/Sydney'),(25,'Europe/Amsterdam'),(26,'Europe/Berlin'),(27,'Europe/Istanbul'),(28,'Europe/London'),(29,'Europe/Madrid'),(30,'Europe/Moscow'),(31,'Europe/Paris'),(32,'Europe/Rome'),(33,'Pacific/Auckland'),(34,'Pacific/Honolulu')

--Email Default Template Insert

INSERT INTO tbl_email_templates(id, template_slug, template_subject, template_message, template_description, module_id, created_on, created_by, is_deleted, is_active,template_name,tenant_id)VALUES (1,'createuser', 'Confirm user registration', '<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">Congratulations! Your SpurtCMS account has been successfully created.</p><p style="color:#000000;font-size:14px;">Log in at</p><p  style="color:#000000;font-size:14px;">{Loginurl}</p><p  style="color:#000000;font-size:14px;margin:0 0 16px;">and start using your Admin Account.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;margin:0 0 12px;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>Spurt CMS Admin</strong></p></td></tr>', 'Change Password: Enable this toggle to activate the Change Password email template, which confirms to users that their password has been successfully changed.', 0, '2023-03-14 11:09:12', 1, 0, 1,'Create Admin User',1);

INSERT INTO tbl_email_templates(id, template_slug, template_subject, template_message, template_description, module_id, created_on, created_by, is_deleted, is_active,template_name,tenant_id)VALUES (2, 'changepassword', 'New password', '<p>Hello,{firstname} your password updated Successfully. your new password:{password}</p>', 'Change Password: Change Password email template, which confirms to users that their password has been successfully changed.', 0, '2023-03-14 11:09:12', 1, 0, 1,'Change Password',1);

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(3,'OTPGenerate','Email OTP','OTP email template, which sends a One-Time Password (OTP) to users for verification.','<tr><td><p style="margin-left:0;"><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  16px 0;">Dear {FirstName}</h1></td></tr><tr><td><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0;">As requsted, here is your One-Time Password (OTP) to log in to SpurtCMS:</p><h3 style="font-size: 20px; line-height:1; font-weight: bold; color: #000000; margin:20px 0;">OTP:{OTP}</h3><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0 0 20px 0;">Please use this OTP within <span style="font-size: 16px; font-weight:bold;">{Expirytime}</span> to complete the login process. If you do not log in within this time frame, you may request a new OTP.</p><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0 0 20px 0;">If you encounter any issue or need assistance, feel free to reach out to our support team.</p><p style="font-size: 14px; line-height: 24px; color: #000000;margin:0 0 20px 0;">Thank you for using <span style="font-size: 16px;">SpurtCMS.</span></p><p style="font-size: 16px; line-height: normal; color: #000000;margin:0 0 15px 0;">Best Regards,</p><p style="font-size: 16px; line-height:normal; color: #000000;margin:0 0 10px 0;"><b>SpurtCMS Team</b></p></td></tr>',0,'2023-11-24 14:56:12',1, 0,1,'OTP Generation',1)

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(4,'Forgot Password','Reset password link','Forgot Password email template, which sends a confirmation and reset password link to users.', '<tr><td style="padding:2.5rem 1.5rem 0rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.25rem;">Dear <span style="font-family: "Lexend", sans-serif;font-size: 1rem;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{FirstName},</span></h2></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin: 0;">You recently requested to reset your password for spurtCMS admin login.</p><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.5rem">Click the link below to reset it.</p></td></tr><tr><td style="padding: 0rem 1.5rem;"><a href="{Passwordurl}" style="color: #2FACD6;text-decoration: underline;">{Passwordurlpath}</a><p class="mb-1" style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;">If you did not request a password reset, please ignore this email.</p></td></tr><tr><td style=" padding: 0rem 1.5rem 2.5rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 400;line-height: .9375rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: .5rem;">Best Regards,</h2><p style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #152027;">  Spurt CMS Admin</p></td></tr></tr>' ,0,'2023-11-24 14:56:12',1, 0,1,'Forgot Password',1)

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(5,'Createmember','Member registration','Notifies member that they have been created as a member on the platform.','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">Were glad to have you as a member at spurtCMS. Your membership has been confirmed.</p><p style="color:#000000;font-size:14px;">Log in at</p><p  style="color:#000000;font-size:14px;">{Loginurl}</p><p  style="color:#000000;font-size:14px;margin:0 0 16px;">and indulge in Member related activities and explore our exclusive content.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;margin:0 0 12px;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>Spurt CMS Admin</strong></p></td></tr>' ,6,'2023-11-24 14:56:12',1, 0,1,'Send Member Login credentials',1)

-- INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name)VALUES(11,'memberactivation','Your Membership Activated on OwnDesk','For the Member to receive notification when they have been registered and activated as a member on the platform. ','<table width="100%" cellpadding="0" cellspacing="0" role="presentation"><tbody><tr><td style="padding:27px 16px 32px;"><table style="margin-inline:auto;max-width:696px;" width="100%" cellpadding="0" cellspacing="0" role="presentation"><tbody> <tr><td style="text-align:center;"><img style="height:66px;width:106px;" src="{AdminLogo}" alt="logo"></td></tr><tr><td><table style="border-bottom:1px solid #EEEEEE;" width="100%" cellpadding="0" cellspacing="0" role="presentation"><tbody><tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">OwnDesk is glad to notify that you have been registered and activated as an Esteemed Member at our OwnDesk Platform.</p><p style="color:#000000;font-size:14px;margin:0 0 16px;">We look forward to a successful and fruitful collaboration with you.</p></td></tr><tr><td><p style="color:#000000;font-size:14px;line-height:normal;margin:0 0 10px;">Best Regards,</p><p style="color:#000000;font-size:14px;font-weight:500; margin:0 0 15px;"><strong style="margin:0 0 16px;">The Team @OwnDesk</strong></p></td></tr></tbody></table></td></tr><tr><td style="padding-top:16px;text-align:center;"><div style="margin-bottom:1rem;"><a style="display:inline-block;height:1.5rem;margin-right:.75rem;width:1.5rem;" href="{FacebookLink}"><img style="max-width:100%;width:auto;" src="{FbLogo}" alt=""> </a><a style="display:inline-block;height:1.5rem;margin-right:.75rem;width:1.5rem;" href="{LinkedinLink}""><img style="max-width:100%;width:auto;" src="{LinkedinLogo}" alt=""> </a><a style="display:inline-block;height:1.5rem;margin-right:.75rem;width:1.5rem;" href="{TwitterLink} "><img style="max-width:100%;width:auto;" src="{TwitterLogo}" alt=""> </a><a style="display:inline-block;height:1.5rem;margin-right:.75rem;width:1.5rem;" href="{YoutubeLink}"><img style="max-width:100%;width:auto;" src="{YoutubeLogo}" alt="InstaLogo"> </a><a style="display:inline-block;height:1.5rem;margin-right:.75rem;width:1.5rem;" href="{InstagramLink}"><img style="max-width:100%;width:auto;" src="{InstaLogo}" alt="InstaLogo"></a></div></td></tr></tbody></table></td></tr></tbody></table>',6,'2023-11-24 14:56:12',1, 0,1,'Send Member Activation Email')

INSERT INTO tbl_email_templates(id, template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(12,'Logined successfully','User login successfully','User conformation account is logined successfully','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">Congratulations! Your SpurtCMS account has been logined successfully.</p><p  style="color:#000000;font-size:14px;margin:0 0 16px;">Start using your Admin Account.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;margin:0 0 12px;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>Spurt CMS Admin</strong></p></td></tr>',0,'2023-11-24 14:56:12',1, 0,1,'Send User Logined Email',1)

INSERT INTO tbl_email_templates(id,template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES(13,'Superadminnotification','New User Registration Alert','For the Super Admin to receive notification when Tenant have been registered as a SpurtCMS platform. ','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>spurtCMS Admin</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">A new user has just registered on spurtCMS. Here are the details:</p></td></tr><tr><td><p style="color:#000000;font-size:14px;"><strong>Name : {FirstName}</strong></p><p  style="color:#000000;font-size:14px;"><strong>Email : {Useremail}</strong></p><p  style="color:#000000;font-size:14px;margin:0 0 16px;"><strong>Registration Time : {Timestamp}</strong></p></td></tr><tr><td><p style="color:#000000;font-size:14px;margin:0 0 12px;">Please review the registration details in the admin panel and take any necessary actions.</p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>The spurtCMS Team</strong></p></td></tr>',0,'2024-11-02 14:56:12',1, 0,1,'New User Registration Alert',1)

-- INSERT INTO tbl_email_templates(template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name)VALUES('SupportEmail','There`s an Inquiry for Support Services','For the Super Admin to receive notification when Tenant have been registered as a Support Services','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>spurtCMS Admin</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;margin:0 0 16px;">You have received an inquiry for Support Services. You may check the details below:</p></td></tr><tr><td><p  style="color:#000000;font-size:14px;margin:0 0 0;"><strong>Service :</strong> {Service}</p><p style="color:#000000;font-size:14px;margin:0 0 0;"><strong>Name :</strong> {FirstName}</p><p  style="color:#000000;font-size:14px;margin:0 0 0;"><strong>Email :</strong> {Useremail}</p><p  style="color:#000000;font-size:14px;margin:0 0 0;"><strong>Time Zone :</strong> {Timestamp}</p><p  style="color:#000000;font-size:14px;margin:0 0 0;"><strong>Number :</strong> {Number}</p><p  style="color:#000000;font-size:14px;margin:0 0 0;"><strong>Country :</strong> {Country}</p><pre style="font-family: `Lexend`, sans-serif;font-size: 14px;text-align: left;color: #000000;margin:0 0 1rem"><strong>Describe :</strong> {Describe}</pre></td></tr><tr><td><p style="color:#000000;font-size:14px;margin:0 0 12px;">Once you have read the details, you may now take the necessary action to provide the required support. </p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:16px 0 0;"><strong>The spurtCMS Team</strong></p></td></tr>',0,'2023-12-24 14:56:12',1, 0,1,'Support Service Email')

-- INSERT INTO tbl_email_templates(template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name)VALUES('SupportuserverificationEmail','Your Inquiry for Support Services Received ','For the User receive notification when Tenant have been registered as a Support Services','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;">We`ve received you inquiry for Support Services.</p></td></tr><tr><td><p style="color:#000000;font-size:14px;margin:0 0 12px;">Your Inquiry is under review and one of our representatives will get back to you on this same email Id shortly.  </p></td></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:16px 0 0;"><strong>The spurtCMS Team</strong></p></td></tr>',0,'2023-12-27 14:56:12',1, 0,1,'Support Service User verification Email')


--Default Page Type

INSERT INTO tbl_page_types(id, title, slug, created_by, created_on, is_deleted) VALUES (1, 'Simple Page', 'simplepage', 1,'2023-03-14 11:09:12',0);

INSERT INTO tbl_page_types(id, title, slug, created_by, created_on, is_deleted) VALUES (2, 'Ecommerce Page', 'ecommercepage', 1,'2023-03-14 11:09:12',0);


-- INSERT INTO tbl_email_templates(template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name)VALUES('registereduserlogin','Registered User Login Alert','For the Super Admin to receive notification when Tenant have been logined as a SpurtCMS platform. ','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 20px; font-weight: bold;line-height: 27px;color:#000000; margin:0 0  12px 0;">Dear <strong>spurtCMS Admin</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:14px;margin:0 0 16px">The User  - {FirstName} has just logged into spurtCMS. Here are the details:</p></td></tr><tr><td><p  style="color:#000000;font-size:14px;"><strong>Registered Email Id : {Useremail}</strong></p><p  style="color:#000000;font-size:14px;margin:0 0 16px;"><strong>Registration Time : {Timestamp}</strong></p></td></tr><tr></tr><tr><td><p style="color:#000000;font-size:16px;line-height:normal;">Best Regards,</p><p style="color:#000000;font-size:16px;font-weight:500;line-height:24px;margin:0 0 16px;"><strong>The spurtCMS Team</strong></p></td></tr>',0,'2024-11-02 14:56:12',1, 0,1,'Registered User Login Alert')



-- Blogs

INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,slug_name) VALUES (1,'Blog', 'blog', 'Blog Templates', 1, 1, 'current_time','blog')

INSERT INTO tbl_templates(id, template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(1,'Contentflow','contentflow_next_js_theme', 1,'blogstarter.png','media/blogstarter.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-starter-theme&demo-title=nextjs-starter-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-starter-theme','current_time',1,1,0,'blog','https://nextjs-starter-theme.vercel.app/')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(2,'Writer`s hub theme','nextjs_writer`s_hub_theme', 1,'blog1.png','media/blog1.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog-theme&demo-title=nextjs-writer`s-hub-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog-theme','current_time',1,1,0,'blog','https://nextjs-blog-theme-liart.vercel.app/')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(3,'Storyze','next-js-blog2-theme', 1,'blog2.png','media/blog2.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog2-theme&demo-title=nextjs-blog2-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog2-theme','current_time',1,1,0,'blog','https://storyze-nextjs-theme.vercel.app/')

INSERT INTO tbl_templates( id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name) VALUES(4,'Content chronicle','content_chronicle_nextJS_theme', 1,'blog4.png','media/blog4.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/nextjs-blog4-theme&demo-title=nextjs-blog4-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-blog4-theme','current_time',1,1,0,'blog')

INSERT INTO tbl_templates( id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name) VALUES(5,'InsightSphere','insightsphere', 1,'blog1.jpg','/media/V1-blog1.jpg','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/v1-blog1-theme&demo-title=v1-blog1-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/v1-blog1-theme','current_time',1,1,0,'blog')

INSERT INTO tbl_templates( id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(6,'ContentVerse','contentverse', 1,'blog2.jpg','/media/V2-blog2.jpg','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/v1-blog2-theme&demo-title=v1-blog2-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/v1-blog2-theme','current_time',1,1,0,'blog','https://content-verse-five.vercel.app/')


-- knowledgeBase

INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,slug_name) VALUES ( 2,'Help Center', 'help-center', 'Help Center Templates', 1, 1, 'current_time','knowledge-base')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(7,'Docuno','docuno-nextjs-theme', 2,'knowledgeBase.png','media/knowledgeBase.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/knowledge-base&demo-title=knowledge-base&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/knowledge-base','current_time',1,1,0,'knowledge-base','https://knowledge-base-lyart-one.vercel.app/')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(8,'Support Sphere','support-sphere', 2,'KnowledgeBase.jpg','/media/KnowledgeBase.jpg','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/help-center&demo-title=knowledge-base&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/help-center','current_time',1,1,0,'knowledge-base','https://support-sphere-gilt.vercel.app/')

-- News

INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,slug_name) VALUES ( 3,'News', 'news', 'News Templates', 1, 1, 'current_time','news')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name) VALUES(9,'Newsprism','newsprism-nextjs-theme', 3,'CoverimageNewsTemplate.jpg','media/CoverimageNewsTemplate.jpg','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/newsprism-nextjs-theme&demo-title=newsprism-nextjs-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/Newsprism-nextjs-theme','current_time',1,1,0,'news')

-- Jobs

INSERT INTO tbl_template_modules(id,template_module_name,template_module_slug,description,is_active,created_by,created_on,slug_name) VALUES (4,'Jobs', 'jobs', 'Jobs Templates', 1, 1, 'current_time','jobs')

INSERT INTO tbl_templates(id,template_name, template_slug, template_module_id, image_name, image_path, deploy_link, github_link, created_on, created_by, is_active, is_deleted,slug_name,preview) VALUES(10,'JobiFylo','jobifylo-nextjs-theme', 4,'jobs.png','media/jobs.png','https://vercel.com/new/clone?repository-url=https://github.com/spurtcms/jobifylo-nextjs-theme&demo-title=jobifylo-nextjs-theme&env=NEXT_PUBLIC_SPURTCMS_NEXTJS_STARTER_APIKEY','https://github.com/spurtcms/nextjs-jobs-theme','current_time',1,1,0,'jobs','https://jobi-fylo-nextjs-theme.vercel.app/')




-- UPDATE tbl_email_templates SET template_message='<tr><td style="padding:2.5rem 1.5rem 0rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.25rem;">Dear <span style="font-family: "Lexend", sans-serif;font-size: 1rem;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{FirstName},</span></h2></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin: 0;">You recently requested to reset your password for spurtCMS admin login.</p><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.5rem">Click the link below to reset it.</p></td></tr><tr><td style="padding: 0rem 1.5rem;"><a href="{Passwordurl}" style="color: #2FACD6;text-decoration: underline;">{Passwordurlpath}</a><p class="mb-1" style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;">If you did not request a password reset, please ignore this email.</p></td></tr><tr><td style=" padding: 0rem 1.5rem 2.5rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 400;line-height: .9375rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: .5rem;">Best Regards,</h2><p style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #152027;">  Spurt CMS Admin</p></td></tr></tr>' WHERE template_slug='Forgot Password'

-- UPDATE tbl_email_templates SET template_subject='Email OTP' where template_slug='OTPGenerate'

  -- paid membership querys start (membership level default templates)


INSERT INTO tbl_mstr_membershiplevels(id,subscription_name,membershiplevel_details,Description,initial_payment,recurrent_subscription,billing_amount,billingfrequent_value,billingfrequent_type,billing_cyclelimit,custom_trial,trial_billing_amount,Trial_billing_limit,created_on,created_by,is_deleted,is_active)Values(1,'Explorer','','Begin your journey with the essentials',0,1,149,4,7,1,0,0,0,'2025-01-25 11:09:12',1,0,1)

INSERT INTO tbl_mstr_membershiplevels(id,subscription_name,membershiplevel_details,Description,initial_payment,recurrent_subscription,billing_amount,billingfrequent_value,billingfrequent_type,billing_cyclelimit,custom_trial,trial_billing_amount,Trial_billing_limit,created_on,created_by,is_deleted,is_active)Values(2,'Insider','','Gain deeper access and start engaging',100,1,249,1,30,6,0,0,0,'2025-01-25 11:09:12',1,0,1)

INSERT INTO tbl_mstr_membershiplevels(id,subscription_name,membershiplevel_details,Description,initial_payment,recurrent_subscription,billing_amount,billingfrequent_value,billingfrequent_type,billing_cyclelimit,custom_trial,trial_billing_amount,Trial_billing_limit,created_on,created_by,is_deleted,is_active)Values(3,'pro','Your Membership Information You have chosen the Pro Membership. $350 / year  The Pro Membership is built for those who want it all — priority access, premium content, and exclusive features that elevate your experience. It`s the perfect choice for professionals and creators who are serious about growth, community, and getting the most value from every interaction. ','For serious users with extended features',100,1,299,6,7,1,0,0,0,'2025-01-25 11:09:12',1,0,1)





-- default country insertion--

INSERT INTO tbl_countries (id, country_code, country_name) VALUES (1, 'AF', 'Afghanistan'), (2, 'AL', 'Albania'), (3, 'DZ', 'Algeria'), (4, 'DS', 'American Samoa'), (5, 'AD', 'Andorra'), (6, 'AO', 'Angola'), (7, 'AI', 'Anguilla'), (8, 'AQ', 'Antarctica'), (9, 'AG', 'Antigua and Barbuda'), (10, 'AR', 'Argentina'), (11, 'AM', 'Armenia'), (12, 'AW', 'Aruba'), (13, 'AU', 'Australia'), (14, 'AT', 'Austria'), (15, 'AZ', 'Azerbaijan'), (16, 'BS', 'Bahamas'), (17, 'BH', 'Bahrain'), (18, 'BD', 'Bangladesh'), (19, 'BB', 'Barbados'), (20, 'BY', 'Belarus'), (21, 'BE', 'Belgium'), (22, 'BZ', 'Belize'), (23, 'BJ', 'Benin'), (24, 'BM', 'Bermuda'), (25, 'BT', 'Bhutan'), (26, 'BO', 'Bolivia'), (27, 'BA', 'Bosnia and Herzegovina'), (28, 'BW', 'Botswana'), (29, 'BV', 'Bouvet Island'), (30, 'BR', 'Brazil'), (31, 'IO', 'British Indian Ocean Territory'), (32, 'BN', 'Brunei Darussalam'), (33, 'BG', 'Bulgaria'), (34, 'BF', 'Burkina Faso'), (35, 'BI', 'Burundi'), (36, 'KH', 'Cambodia'), (37, 'CM', 'Cameroon'), (38, 'CA', 'Canada'), (39, 'CV', 'Cape Verde'), (40, 'KY', 'Cayman Islands'), (41, 'CF', 'Central African Republic'), (42, 'TD', 'Chad'), (43, 'CL', 'Chile'), (44, 'CN', 'China'), (45, 'CX', 'Christmas Island'), (46, 'CC', 'Cocos (Keeling) Islands'), (47, 'CO', 'Colombia'), (48, 'KM', 'Comoros'), (49, 'CD', 'Democratic Republic of the Congo'), (50, 'CG', 'Republic of Congo'), (51, 'CK', 'Cook Islands'), (52, 'CR', 'Costa Rica'), (53, 'HR', 'Croatia (Hrvatska)'), (54, 'CU', 'Cuba'), (55, 'CY', 'Cyprus'), (56, 'CZ', 'Czech Republic'), (57, 'DK', 'Denmark'), (58, 'DJ', 'Djibouti'), (59, 'DM', 'Dominica'), (60, 'DO', 'Dominican Republic'), (61, 'TP', 'East Timor'), (62, 'EC', 'Ecuador'), (63, 'EG', 'Egypt'), (64, 'SV', 'El Salvador'), (65, 'GQ', 'Equatorial Guinea'), (66, 'ER', 'Eritrea'), (67, 'EE', 'Estonia'), (68, 'ET', 'Ethiopia'), (69, 'FK', 'Falkland Islands (Malvinas)'), (70, 'FO', 'Faroe Islands'), (71, 'FJ', 'Fiji'), (72, 'FI', 'Finland'), (73, 'FR', 'France'), (74, 'FX', 'France, Metropolitan'), (75, 'GF', 'French Guiana'), (76, 'PF', 'French Polynesia'), (77, 'TF', 'French Southern Territories'), (78, 'GA', 'Gabon'), (79, 'GM', 'Gambia'), (80, 'GE', 'Georgia'), (81, 'DE', 'Germany'), (82, 'GH', 'Ghana'), (83, 'GI', 'Gibraltar'), (84, 'GK', 'Guernsey'), (85, 'GR', 'Greece'), (86, 'GL', 'Greenland'), (87, 'GD', 'Grenada'), (88, 'GP', 'Guadeloupe'), (89, 'GU', 'Guam'), (90, 'GT', 'Guatemala'), (91, 'GN', 'Guinea'), (92, 'GW', 'Guinea-Bissau'), (93, 'GY', 'Guyana'), (94, 'HT', 'Haiti'), (95, 'HM', 'Heard and Mc Donald Islands'), (96, 'HN', 'Honduras'), (97, 'HK', 'Hong Kong'), (98, 'HU', 'Hungary'), (99, 'IS', 'Iceland'), (100, 'IN', 'India'), (101, 'IM', 'Isle of Man'), (102, 'ID', 'Indonesia'), (103, 'IR', 'Iran (Islamic Republic of)'), (104, 'IQ', 'Iraq'), (105, 'IE', 'Ireland'), (106, 'IL', 'Israel'), (107, 'IT', 'Italy'), (108, 'CI', 'Ivory Coast'), (109, 'JE', 'Jersey'), (110, 'JM', 'Jamaica'), (111, 'JP', 'Japan'), (112, 'JO', 'Jordan'), (113, 'KZ', 'Kazakhstan'), (114, 'KE', 'Kenya'), (115, 'KI', 'Kiribati'), (116, 'KP', 'Korea, Democratic People`s Republic of'), (117, 'KR', 'Korea, Republic of'), (118, 'XK', 'Kosovo'), (119, 'KW', 'Kuwait'), (120, 'KG', 'Kyrgyzstan'), (121, 'LA', 'Lao People`s Democratic Republic'), (122, 'LV', 'Latvia'), (123, 'LB', 'Lebanon'), (124, 'LS', 'Lesotho'), (125, 'LR', 'Liberia'), (126, 'LY', 'Libyan Arab Jamahiriya'), (127, 'LI', 'Liechtenstein'), (128, 'LT', 'Lithuania'), (129, 'LU', 'Luxembourg'), (130, 'MO', 'Macau'), (131, 'MK', 'North Macedonia'), (132, 'MG', 'Madagascar'), (133, 'MW', 'Malawi'), (134, 'MY', 'Malaysia'), (135, 'MV', 'Maldives'), (136, 'ML', 'Mali'), (137, 'MT', 'Malta'), (138, 'MH', 'Marshall Islands'), (139, 'MQ', 'Martinique'), (140, 'MR', 'Mauritania'), (141, 'MU', 'Mauritius'), (142, 'TY', 'Mayotte'), (143, 'MX', 'Mexico'), (144, 'FM', 'Micronesia, Federated States of'), (145, 'MD', 'Moldova, Republic of'), (146, 'MC', 'Monaco'), (147, 'MN', 'Mongolia'), (148, 'ME', 'Montenegro'), (149, 'MS', 'Montserrat'), (150, 'MA', 'Morocco'), (151, 'MZ', 'Mozambique'), (152, 'MM', 'Myanmar'), (153, 'NA', 'Namibia'), (154, 'NR', 'Nauru'), (155, 'NP', 'Nepal'), (156, 'NL', 'Netherlands'), (157, 'AN', 'Netherlands Antilles'), (158, 'NC', 'New Caledonia'), (159, 'NZ', 'New Zealand'), (160, 'NI', 'Nicaragua'), (161, 'NE', 'Niger'), (162, 'NG', 'Nigeria'), (163, 'NU', 'Niue'), (164, 'NF', 'Norfolk Island'), (165, 'MP', 'Northern Mariana Islands'), (166, 'NO', 'Norway'), (167, 'OM', 'Oman'), (168, 'PK', 'Pakistan'), (169, 'PW', 'Palau'), (170, 'PS', 'Palestine'), (171, 'PA', 'Panama'), (172, 'PG', 'Papua New Guinea'), (173, 'PY', 'Paraguay'), (174, 'PE', 'Peru'), (175, 'PH', 'Philippines'), (176, 'PN', 'Pitcairn'), (177, 'PL', 'Poland'), (178, 'PT', 'Portugal'), (179, 'PR', 'Puerto Rico'), (180, 'QA', 'Qatar'), (181, 'RE', 'Reunion'), (182, 'RO', 'Romania'), (183, 'RU', 'Russian Federation'), (184, 'RW', 'Rwanda'), (185, 'KN', 'Saint Kitts and Nevis'), (186, 'LC', 'Saint Lucia'), (187, 'VC', 'Saint Vincent and the Grenadines'), (188, 'WS', 'Samoa'), (189, 'SM', 'San Marino'), (190, 'ST', 'Sao Tome and Principe'), (191, 'SA', 'Saudi Arabia'), (192, 'SN', 'Senegal'), (193, 'RS', 'Serbia'), (194, 'SC', 'Seychelles'), (195, 'SL', 'Sierra Leone'), (196, 'SG', 'Singapore'), (197, 'SK', 'Slovakia'), (198, 'SI', 'Slovenia'), (199, 'SB', 'Solomon Islands'), (200, 'SO', 'Somalia'), (201, 'ZA', 'South Africa'), (202, 'GS', 'South Georgia South Sandwich Islands'), (203, 'SS', 'South Sudan'), (204, 'ES', 'Spain'), (205, 'LK', 'Sri Lanka'), (206, 'SH', 'St. Helena'), (207, 'PM', 'St. Pierre and Miquelon'), (208, 'SD', 'Sudan'), (209, 'SR', 'Suriname'), (210, 'SJ', 'Svalbard and Jan Mayen Islands'), (211, 'SZ', 'Swaziland'), (212, 'SE', 'Sweden'), (213, 'CH', 'Switzerland'), (214, 'SY', 'Syrian Arab Republic'), (215, 'TW', 'Taiwan'), (216, 'TJ', 'Tajikistan'), (217, 'TZ', 'Tanzania, United Republic of'), (218, 'TH', 'Thailand'), (219, 'TG', 'Togo'), (220, 'TK', 'Tokelau'), (221, 'TO', 'Tonga'), (222, 'TT', 'Trinidad and Tobago'), (223, 'TN', 'Tunisia'), (224, 'TR', 'Turkey'), (225, 'TM', 'Turkmenistan'), (226, 'TC', 'Turks and Caicos Islands'), (227, 'TV', 'Tuvalu'), (228, 'UG', 'Uganda'), (229, 'UA', 'Ukraine'), (230, 'AE', 'United Arab Emirates'), (231, 'GB', 'United Kingdom'), (232, 'US', 'United States'), (233, 'UM', 'United States minor outlying islands'), (234, 'UY', 'Uruguay'), (235, 'UZ', 'Uzbekistan'), (236, 'VU', 'Vanuatu'), (237, 'VA', 'Vatican City State'), (238, 'VE', 'Venezuela'), (239, 'VN', 'Vietnam'), (240, 'VG', 'Virgin Islands (British)'), (241, 'VI', 'Virgin Islands (U.S.)'), (242, 'WF', 'Wallis and Futuna Islands'), (243, 'EH', 'Western Sahara'), (244, 'YE', 'Yemen'), (245, 'ZM', 'Zambia'), (246, 'ZW', 'Zimbabwe');


-- INSERT INTO tbl_go_templates (id,template_name,template_image,is_deleted,channel_slug_name,template_module_name,created_on) VALUES (1,'Content chronicle','/image-resize?name=media/blog4.png',0,'blog','Blog','current_time')

-- INSERT INTO tbl_go_templates (id,template_name,template_image,is_deleted,channel_slug_name,template_module_name,created_on) VALUES (2,'Content flow','/image-resize?name=media/blogstarter.png',0,'blog','Blog','current_time')

-- INSERT INTO tbl_go_templates (id,template_name,template_image,is_deleted,channel_slug_name,template_module_name,created_on) VALUES (3,'Writer`s hub theme','/image-resize?name=media/blog1.png',0,'blog','Blog','current_time')

-- INSERT INTO tbl_go_templates (id,template_name,template_image,is_deleted,channel_slug_name,template_module_name,created_on) VALUES (4,'Content Verse','/image-resize?name=/media/V2-blog2.jpg',0,'blog','Blog','current_time')

-- INSERT INTO tbl_go_templates (id,template_name,template_image,is_deleted,channel_slug_name,template_module_name,created_on) VALUES (5,'Support Sphere','/image-resize?name=/media/KnowledgeBase.jpg',0,'knowledge-base','Help Center','current_time')

-- INSERT INTO tbl_go_templates (id,template_name,template_image,is_deleted,channel_slug_name,template_module_name,created_on) VALUES (6,'Courses','/image-resize?name=/media/courses.jpg',0,'course','Courses','current_time')


-- UPDATE tbl_go_templates SET created_on='2025-08-22 10:45:21.109783' where tenant_id is Null

-- truncate table tbl_go_templates

INSERT INTO tbl_go_templates (id,template_name,template_image,is_deleted,channel_slug_name,template_module_name,created_on,template_description,tenant_id,created_by) VALUES (1,'Content flow','/public/img/content_flow.png',0,'blog','Blog','current_time','A content flow template visually maps out each stage of content creation, from planning to publication.It ensures a structured, repeatable process for teams to follow and maintain content quality.','1',1)
-- truncate table tbl_websites
INSERT INTO tbl_websites(id,name,channel_names,template_id,tenant_id,is_deleted,created_by,created_on)values(1,'base_url','Blog',1,'1',0,1,'current_time')

INSERT INTO tbl_menus(id,name,description,slug_name, status,parent_id,tenant_id,created_on,created_by,is_deleted,website_id)VALUES(1,'Header','The primary menu at the top of a webpage, guiding users to major site sections','header',1,0,'1','current_time', 1,0,1)