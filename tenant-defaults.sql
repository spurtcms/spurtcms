-- Must use separate insert queries for parent and child data insertions to maintain sequence.

-- Some reserved words are used for dynamic values insertion in queries.They are given as below.Use them whenever they needed essentially.

-- current-time,u_id(user id),'tid'(tenant id),pid(parent id),chid(channel id),blid(block id),tagid(tag id),mapcat(mapped categories)

INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Default Category', 'default_category', 'current-time', u_id, 0, 0, 'Default_Category', 'tid');

INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Default1', 'default1', 'current-time', u_id, 0, pid, 'Default1','tid');




--  default Parent blog CATEGORY for channel based templates
-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by, is_deleted, parent_id, description, tenant_id) VALUES ('Blog', 'blog', 'current-time', u_id, 0,0,'Parent Category for blogs', 'tid')

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Technology', 'technology', 'current-time', u_id, 0, pid, 'Technology','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Lifestyle', 'lifestyle', 'current-time', u_id, 0, pid, 'Lifestyle','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Travel', 'travel', 'current-time', u_id, 0, pid, 'Travel','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Business', 'business', 'current-time', u_id, 0, pid, 'Business','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Education', 'education', 'current-time', u_id, 0, pid, 'Education','tid');


-- Ecommerce


-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Ecommerce', 'ecommerce', 'current-time', u_id, 0, 0, 'Ecommerce', 'tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Fashion & Clothing', 'fashion_clothing', 'current-time', u_id, 0, sp_id, 'Fashion & Clothing','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Men`s Clothing', 'mens_clothing', 'current-time', u_id, 0, sub_id, 'Men`s Clothing','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Women`s Clothing', 'womens_clothing', 'current-time', u_id, 0, sub_id, 'Women`s Clothing','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Kid`s Clothing', 'sids_clothing', 'current-time', u_id, 0, sub_id, 'Kid`s Clothing','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Accessories', 'accessories', 'current-time', u_id, 0, sub_id, 'Accessories','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Footwear', 'footwear', 'current-time', u_id, 0, sub_id, 'Footwear','tid');


-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Electronics', 'electronics', 'current-time', u_id, 0, sp_id, 'Electronics','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Mobile Phones', 'mobile_phones', 'current-time', u_id, 0, e_id, 'Mobile Phones','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Laptops & Computers', 'laptops_computers', 'current-time', u_id, 0, e_id, 'Laptops & Computers','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Cameras', 'cameras', 'current-time', u_id, 0, e_id, 'Cameras','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Home Appliances', 'home_appliances', 'current-time', u_id, 0, e_id, 'Home Appliances','tid');


-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Home & Furniture', 'home_furniture', 'current-time', u_id, 0, sp_id, 'Home & Furniture','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Living Room', 'living_room', 'current-time', u_id, 0, h_id, 'Living Room','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Bedroom', 'bedroom', 'current-time', u_id, 0, h_id, 'Bedroom','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Kitchen', 'kitchen', 'current-time', u_id, 0, h_id, 'Kitchen','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Decor', 'decor', 'current-time', u_id, 0, h_id, 'Decor','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Beauty & Personal Care', 'beauty_personal_care', 'current-time', u_id, 0, sp_id, 'Beauty & Personal Care','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Makeup', 'makeup', 'current-time', u_id, 0, b_id, 'Makeup','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Skincare', 'skincare', 'current-time', u_id, 0, b_id, 'Skincare','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Haircare', 'haircare', 'current-time', u_id, 0, b_id, 'Haircare','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Fragrances', 'fragrances', 'current-time', u_id, 0, b_id, 'Fragrances','tid');



-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Sports & Outdoors', 'sports_outdoors', 'current-time', u_id, 0, sp_id, 'Sports & Outdoors','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Fitness Equipment', 'fitness_equipment', 'current-time', u_id, 0, so_id, 'Fitness Equipment','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Sports Gear', 'sports_gear', 'current-time', u_id, 0, so_id, 'Sports Gear','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Activewear', 'activewear', 'current-time', u_id, 0, so_id, 'Activewear','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Outdoor', 'outdoor', 'current-time', u_id, 0, so_id, 'Outdoor','tid');


-- -- jobs

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Jobs', 'jobs', 'current-time', u_id, 0, 0, 'Jobs', 'tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Healthcare', 'healthcare', 'current-time', u_id, 0, pid, 'Healthcare','tid');
-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Qualification', 'qualification', 'current-time', u_id, 0, pid, 'Education','tid');
-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Finance', 'finance', 'current-time', u_id, 0, pid, 'Finance','tid');
-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Marketing & Sales', 'marketing_sales', 'current-time', u_id, 0, pid, 'Marketing & Sales','tid');
-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Engineering', 'engineering', 'current-time', u_id, 0, pid, 'Engineering','tid');

-- -- news

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('News', 'news', 'current-time', u_id, 0, 0, 'News', 'tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Politics', 'politics', 'current-time', u_id, 0, pid, 'Politics','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('National', 'national', 'current-time', u_id, 0, po_id, 'National','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('International', 'international', 'current-time', u_id, 0, po_id, 'International','tid');


-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Finance', 'finance', 'current-time', u_id, 0, pid, 'Finance','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Markets', 'markets', 'current-time', u_id, 0, bc_id, 'Markets','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Economy', 'economy', 'current-time', u_id, 0, bc_id, 'Economy','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Corporate', 'corporate', 'current-time', u_id, 0, bc_id, 'Corporate','tid');


-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Tech Insights', 'tech_insights', 'current-time', u_id, 0, pid, 'Tech Insights','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Gadgets', 'gadgets', 'current-time', u_id, 0, tc_id, 'Gadgets','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Innovation', 'innovation', 'current-time', u_id, 0, tc_id, 'Innovation','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Apps & Software', 'apps_software', 'current-time', u_id, 0, tc_id, 'Apps & Software','tid');


-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Health', 'health', 'current-time', u_id, 0, pid, 'Health','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Physical Health', 'physical_health', 'current-time', u_id, 0, hc_id, 'Physical Health','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Mental Health', 'mental_health', 'current-time', u_id, 0, hc_id, 'Mental Health','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Healthcare Industry', 'healthcare_industry', 'current-time', u_id, 0, hc_id, 'Healthcare Industry','tid');


-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Sports', 'sports', 'current-time', u_id, 0, pid, 'Sports','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Cricket', 'cricket', 'current-time', u_id, 0, sc_id, 'Cricket','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Football', 'football', 'current-time', u_id, 0, sc_id, 'Football','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Other Sports', 'other_sports', 'current-time', u_id, 0, sc_id, 'Other Sports','tid');


-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('World News', 'worldnews', 'current-time', u_id, 0, pid, 'World News','tid');

--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Asia', 'asia', 'current-time', u_id, 0, w_id, 'Asia','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Middle east', 'Middle East', 'current-time', u_id, 0, w_id, 'Middle East','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('North america', 'North America', 'current-time', u_id, 0, w_id, 'North America','tid');
--     INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Europe', 'europe', 'current-time', u_id, 0, w_id, 'Europe','tid');



-- -- ****************** Knowledge-Base ***********************
-- --  default knowledge base categories for channel based templates
-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by, is_deleted, parent_id, description, tenant_id) VALUES ('Knowledge Base', 'knowledge-base', 'current-time', u_id, 0,0,'Parent Category for knowledge base', 'tid')

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Getting Started', 'gettingstarted', 'current-time', u_id, 0, pid, 'Getting Started','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Product Documentation', 'productdocumentation', 'current-time', u_id, 0, pid, 'Product Documentation','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Troubleshooting', 'troubleshooting', 'current-time', u_id, 0, pid, 'Troubleshooting','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Billing & Payments', 'billingpayments', 'current-time', u_id, 0, pid, 'Billing & Payments','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Integrations', 'Integrations', 'current-time', u_id, 0, pid, 'Integrations','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Release Notes', 'release_notes', 'current-time', u_id, 0, pid, 'Release Notes','tid');

--Default Courses Category

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by, is_deleted, parent_id, description, tenant_id) VALUES ('Courses', 'courses', 'current-time', u_id, 0,0,'Parent Category for Courses', 'tid')

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Science & Technology', 'science-technology', 'current-time', u_id, 0, pid, 'Science & Technology','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Business & Management', 'business-management', 'current-time', u_id, 0, pid, 'Business & Management','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Arts & Humanities', 'arts-humanities', 'current-time', u_id, 0, pid, 'Arts & Humanities','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Engineering', 'engineering', 'current-time', u_id, 0, pid, 'Engineering','tid');

-- INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Computer Science & IT', 'computerscience-it', 'current-time', u_id, 0, pid, 'Computer Science & IT','tid');


-- channels default 

INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id,channel_unique_id) VALUES ('Default Channel', 'default_channel', 0, 1, 0, 'current-time', u_id, 'default description','mychannels','tid','duid');

INSERT INTO tbl_channel_categories(channel_id, category_id, created_at, created_on,tenant_id) VALUES (chid,'mapcat', u_id, 'current-time','tid');

-- INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('Appointment', 'appointment', 0, 1, 0, 'current-time', u_id, 'Appointment description','mychannels','tid');


-- INSERT INTO tbl_fields(field_name,field_type_id,order_index,image_path,created_on,created_by,is_deleted,tenant_id)VALUES('Date',6,1,'/public/img/date.svg','current-time', u_id,0,'tid')

-- INSERT INTO tbl_fields(field_name,field_type_id,order_index,image_path,created_on,created_by,is_deleted,tenant_id)VALUES('Time',4,2,'/public/img/date-time.svg','current-time', u_id,0,'tid')

-- INSERT INTO tbl_fields(field_name,field_type_id,order_index,image_path,created_on,created_by,is_deleted,tenant_id)VALUES('contact number',2,3,'/public/img/text.svg','current-time', u_id,0,'tid')

-- INSERT INTO tbl_fields(field_name,field_type_id,order_index,image_path,created_on,created_by,is_deleted,tenant_id)VALUES('Email',2,4,'/public/img/text.svg','current-time', u_id,0,'tid')

-- INSERT INTO tbl_fields(field_name,field_type_id,order_index,image_path,created_on,created_by,is_deleted,tenant_id)VALUES('DOB',6,5,'/public/img/date.svg','current-time', u_id,0,'tid')


-- channel tbl_group_fields

-- INSERT INTO tbl_group_fields(channel_id,field_id,tenant_id)VALUES(19,1,'tid')
-- INSERT INTO tbl_group_fields(channel_id,field_id,tenant_id)VALUES(19,2,'tid')
-- INSERT INTO tbl_group_fields(channel_id,field_id,tenant_id)VALUES(19,3,'tid')
-- INSERT INTO tbl_group_fields(channel_id,field_id,tenant_id)VALUES(19,4,'tid')
-- INSERT INTO tbl_group_fields(channel_id,field_id,tenant_id)VALUES(19,5,'tid')


INSERT INTO tbl_member_groups(name,slug,description,is_active,is_deleted,created_on,created_by,tenant_id) VALUES ('Default Group', 'default-group', 'Default Group', 1,0, 'current-time', u_id,'tid');

--  only three block is tenant based


INSERT INTO tbl_member_settings (allow_registration,member_login,notification_users,tenant_id) VALUES (1,'password','u_id','tid')

INSERT INTO tbl_email_configurations (selected_type,tenant_id) VALUES ('environment','tid')

INSERT INTO tbl_general_settings (date_format,expand_logo_path,language_id,logo_path,storage_type,tenant_id,time_format,time_zone) VALUES ('dd mmm yyyy','/public/img/logo-bg.svg',1,'/public/img/logo1.svg','aws','tid','12','Asia/Kolkata')

INSERT INTO tbl_graphql_settings (token_name,description,duration,created_by,created_on,is_deleted,token,is_default,tenant_id) VALUES ('Default Token','Default Token','Unlimited',u_id,'current-time',0,'apikey',1,'tid')

-- below insert queries are for CREATING DEFAULT CATEGORIES, CHANNELS and TEMPLATES


-- default knowledge base channels for channel based templates
-- INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('Knowledge base', 'knowledge-base', 0, 1, 0, 'current-time', u_id, 'Centralizes informative content and resources as Entries. Each entry will go into categories like Tutorials, Guides, or Best Practices for easy reference. Each entry is mapped to a Knowledge Template, ensuring clarity and a professional layout for readers.','mychannels','tid');

-- INSERT INTO tbl_module_permissions(route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES ('rname','knowledge-base','knowledge-base',mid,1,0,1,2,u_id,'current-time','tid')

-- end 


-- default blog and channel based templates
-- INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('Blog', 'blog', 0, 1, 0, 'current-time', u_id, 'Showcases blog posts as Entries, grouped into categories like Lifestyle, Technology, or Travel. Each blog uChannel for blog templatesses a Blog Template for engaging and readable layouts.','mychannels','tid');

-- INSERT INTO tbl_module_permissions(route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES ('rname','Blog','blog',mid,1,0,1,2,u_id,'current-time','tid')

-- end

-- default news and channel based templates

-- INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('News', 'news', 0, 1, 0, 'current-time', u_id, 'Organizes and presents news articles as Entries. News articles can be grouped into categories like Politics, Sports, or Technology for effortless browsing. Each article is linked to a News Template, providing an engaging display with headlines, images, and detailed content.','mychannels','tid');

-- INSERT INTO tbl_module_permissions(route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES ('rname','News','news',mid,1,0,1,2,u_id,'current-time','tid')

-- end 


-- default jobs and channel based templates

-- INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('Jobs', 'jobs', 0, 1, 0, 'current-time', u_id, 'Organizes and displays Job Descriptions (JDs) as Entries. JDs can belong to various categories like Technical, Sales, or Administration for easier navigation. Each JD is linked to a Job Description Template, showcasing key roles, responsibilities, required skills, and more.','mychannels','tid');

-- INSERT INTO tbl_module_permissions(route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES ('rname','Jobs','jobs',mid,1,0,1,2,u_id,'current-time','tid')

-- end

-- default ecommerce and channel based templates

-- INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('Ecommerce', 'ecommerce', 0, 1, 0, 'current-time', u_id, 'Organizes and showcases products as Entries. Products can be grouped into categories like Electronics, Apparel, or Home Decor for seamless navigation. Each product is linked to a Product Listing Template, highlighting specifications, features, pricing, and purchase options.','mychannels','tid');
-- INSERT INTO tbl_module_permissions(route_name,display_name,slug_name,module_id,full_access_permission,parent_id,assign_permission,order_index,created_by,created_on,tenant_id) VALUES ('rname','Ecommerce','ecommerce',mid,1,0,1,2,u_id,'current-time','tid')

-- end

-- channel fields 
-- INSERT INTO tbl_channel_categories(channel_id, category_id, created_at, created_on,tenant_id) VALUES (chid,'mapcat', u_id, 'current-time','tid');


-- default Email-template based on tenant

INSERT INTO tbl_email_templates(template_slug, template_subject, template_message, template_description, module_id, created_on, created_by, is_deleted, is_active,template_name,tenant_id)VALUES ('createuser', 'Confirm user registration', '<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 1.25rem; font-weight: bold;line-height: 1.6875rem;color:#000000; margin:0 0  .75rem 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:.875rem;">Congratulations! Your SpurtCMS account has been successfully created.</p><p style="color:#000000;font-size:.875rem;">Log in at</p><p  style="color:#000000;font-size:.875rem;">{Loginurl}</p><p  style="color:#000000;font-size:.875rem;margin:0 0 1rem;">and start using your Admin Account.</p></td></tr><tr><td><p style="color:#000000;font-size:1rem;line-height:normal;margin:0 0 .75rem;">Best Regards,</p><p style="color:#000000;font-size:1rem;font-weight:500;line-height:1.5rem;margin:0 0 1rem;"><strong>Spurt CMS Admin</strong></p></td></tr>', 'Change Password: Enable this toggle to activate the Change Password email template, which confirms to users that their password has been successfully changed.', 0, 'current-time', 1, 0, 1,'Create Admin User','tid');

INSERT INTO tbl_email_templates(template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES('OTPGenerate','Email OTP','OTP email template, which sends a One-Time Password (OTP) to users for verification.','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 1.25rem; font-weight: bold;line-height: 1.6875rem;color:#000000; margin:0 0  1rem 0;">Dear {FirstName}</h1></td></tr><tr><td><p style="font-size: .875rem; line-height: 1.5rem; color: #000000;margin:0;">As requsted, here is your One-Time Password (OTP) to log in to SpurtCMS:</p><h3 style="font-size: 1.25rem; line-height:1; font-weight: bold; color: #000000; margin:1.25rem 0;">OTP:{OTP}</h3><p style="font-size: .875rem; line-height: 1.5rem; color: #000000;margin:0 0 1.25rem 0;">Please use this OTP within <span style="font-size: 1rem; font-weight:bold;">{Expirytime}</span> to complete the login process. If you do not log in within this time frame, you may request a new OTP.</p><p style="font-size: .875rem; line-height: 1.5rem; color: #000000;margin:0 0 1.25rem 0;">If you encounter any issue or need assistance, feel free to reach out to our support team.</p><p style="font-size: .875rem; line-height: 1.5rem; color: #000000;margin:0 0 1.25rem 0;">Thank you for using <span style="font-size: 1rem;">SpurtCMS.</span></p><p style="font-size: 1rem; line-height: normal; color: #000000;margin:0 0 .9375rem 0;">Best Regards,</p><p style="font-size: 1rem; line-height:normal; color: #000000;margin:0 0 .625rem 0;"><b>SpurtCMS Team</b></p></td></tr>',0,'current-time',1, 0,1,'OTP Generation','tid')

INSERT INTO tbl_email_templates(template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES('Createmember','Member registration','Notifies member that they have been created as a member on the platform.','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 1.25rem; font-weight: bold;line-height: 1.6875rem;color:#000000; margin:0 0  .75rem 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:.875rem;">Were glad to have you as a member at spurtCMS. Your membership has been confirmed.</p><p style="color:#000000;font-size:.875rem;">Log in at</p><p  style="color:#000000;font-size:.875rem;">{Loginurl}</p><p  style="color:#000000;font-size:.875rem;margin:0 0 1rem;">and indulge in Member related activities and explore our exclusive content.</p></td></tr><tr><td><p style="color:#000000;font-size:1rem;line-height:normal;margin:0 0 .75rem;">Best Regards,</p><p style="color:#000000;font-size:1rem;font-weight:500;line-height:1.5rem;margin:0 0 1rem;"><strong>Spurt CMS Admin</strong></p></td></tr>' ,6,'current-time',1, 0,1,'Send Member Login credentials','tid')

INSERT INTO tbl_email_templates(template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES('Logined successfully','User login successfully','User conformation account is login successfully','<tr><td><p style="margin-left:0;">&nbsp;</p><h1 style="font-size: 1.25rem; font-weight: bold;line-height: 1.6875rem;color:#000000; margin:0 0  .75rem 0;">Dear <strong>{FirstName}</strong>,</h1></td></tr><tr><td><p style="color:#000000;font-size:.875rem;">Congratulations! Your SpurtCMS account has been logged in successfully.</p><p  style="color:#000000;font-size:.875rem;margin:0 0 1rem;">Start using your Admin Account.</p></td></tr><tr><td><p style="color:#000000;font-size:1rem;line-height:normal;margin:0 0 .75rem;">Best Regards,</p><p style="color:#000000;font-size:1rem;font-weight:500;line-height:1.5rem;margin:0 0 1rem;"><strong>Spurt CMS Admin</strong></p></td></tr>',0,'current-time',1, 0,1,'Send User Login Email','tid')

INSERT INTO tbl_email_templates(template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES('Forgot Password','Reset password link','Forgot Password email template, which sends a confirmation and reset password link to users.', '<tr><td style="padding: 1.7188rem 1.5rem;border-bottom: .0625rem solid #F5F5F6;background-color: #FDFDFE;"><a href="javascript:void(0)" style="display: inline-block;"><img src="{AdminLogo}" alt="" style="display: block; height:20px"></a></td></tr><tr><td style="padding:2.5rem 1.5rem 0rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.25rem;">Dear <span style="font-family: "Lexend", sans-serif;font-size: 1rem;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{FirstName},</span></h2></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin: 0;">You recently requested to reset your password for spurtCMS admin login.</p><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.5rem">Click the link below to reset it.</p></td></tr><tr><td style="padding: 0rem 1.5rem;"><a href="{Passwordurl}" style="color: #2FACD6;text-decoration: underline;">{Passwordurlpath}</a><p class="mb-1" style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;">If you did not request a password reset, please ignore this email.</p></td></tr><tr><td style=" padding: 0rem 1.5rem 2.5rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 400;line-height: .9375rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: .5rem;">Best Regards,</h2><p style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #152027;">  Spurt CMS Admin</p></td></tr><tr><td style="padding: 1.5rem 1.5rem 2.5rem;text-align: center;border-top: 1px solid #EDEDED;"><a href="javascript:void(0)" style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 300;line-height: .9375rem;letter-spacing: 0em;text-align: center;color: #525252;text-decoration: none;">Privacy policy • Contact us • Read our blog • Join SpurtCMS</a><div style="margin-top: 1rem;"><a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{TwitterLogo}" alt="" style="margin-right: 1rem;"></a> <a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{FbLogo}" alt=""    style="margin-right: 1rem;"></a><a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{LinkedinLogo}" alt="" style="margin-right: 1rem;"></a></div></td></tr></tr>' ,6,'2023-11-24 14:56:12',1, 0,1,'Forgot Password','tid')

INSERT INTO tbl_email_templates(template_slug,template_subject,template_description,template_message,module_id,created_on,created_by,is_deleted,is_active,template_name,tenant_id)VALUES('Forgot Password','Reset password link','Forgot Password email template, which sends a confirmation and reset password link to users.', '<tr><td style="padding: 1.7188rem 1.5rem;border-bottom: .0625rem solid #F5F5F6;background-color: #FDFDFE;"><a href="javascript:void(0)" style="display: inline-block;"><img src="{AdminLogo}" alt="" style="display: block; height:20px"></a></td></tr><tr><td style="padding:2.5rem 1.5rem 0rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.25rem;">Dear <span style="font-family: "Lexend", sans-serif;font-size: 1rem;font-weight: 500;line-height: 20px;letter-spacing: 0em;text-align: left;color: #152027;">{FirstName},</span></h2></td></tr><tr><td style="padding: 0rem 1.5rem;"><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin: 0;">You recently requested to reset your password for spurtCMS admin login.</p><p style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: 1.5rem">Click the link below to reset it.</p></td></tr><tr><td style="padding: 0rem 1.5rem;"><a href="{Passwordurl}" style="color: #2FACD6;text-decoration: underline;">{Passwordurlpath}</a><p class="mb-1" style="font-family: "Lexend", sans-serif;font-size: 14px;font-weight: 400;line-height: 18px;letter-spacing: 0em;text-align: left;color: #525252;">If you did not request a password reset, please ignore this email.</p></td></tr><tr><td style=" padding: 0rem 1.5rem 2.5rem;"><h2 style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 400;line-height: .9375rem;letter-spacing: 0em;text-align: left;color: #525252;margin-bottom: .5rem;">Best Regards,</h2><p style="font-family: "Lexend", sans-serif;font-size: .875rem;font-weight: 400;line-height: 1.125rem;letter-spacing: 0em;text-align: left;color: #152027;">  Spurt CMS Admin</p></td></tr><tr><td style="padding: 1.5rem 1.5rem 2.5rem;text-align: center;border-top: 1px solid #EDEDED;"><a href="javascript:void(0)" style="font-family: "Lexend", sans-serif;font-size: .75rem;font-weight: 300;line-height: .9375rem;letter-spacing: 0em;text-align: center;color: #525252;text-decoration: none;">Privacy policy • Contact us • Read our blog • Join SpurtCMS</a><div style="margin-top: 1rem;"><a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{TwitterLogo}" alt="" style="margin-right: 1rem;"></a> <a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{FbLogo}" alt=""    style="margin-right: 1rem;"></a><a href="javascript:void(0)" style="display: inline-block;"><img style="width:24px" src="{LinkedinLogo}" alt="" style="margin-right: 1rem;"></a></div></td></tr></tr>' ,6,'2023-11-24 14:56:12',1, 0,1,'Forgot Password','tid')


 INSERT INTO tbl_forms(form_title,form_slug,form_data,status,is_active,created_by,created_on,is_deleted,uuid,form_image_path,form_description,channel_name,tenant_id,modified_on,channel_id,form_preview_imagepath)VALUES('Connect with Us','connect-with-us','{"data":[{"uniqID":"98c2d620-41fa-446b-bcf4-5853dabeeba9","type":"Title","value":"Connect with Us"},{"type":"ShortAnswer","uniqID":"649fd04f-79ad-4150-a8ce-00032798f53f","question":"Name","placeholder":"Enter Your Name","value":"","required":false,"min":"","max":"","textareaHeight":"4rem"},{"type":"ShortAnswer","uniqID":"9bc4d442-17bb-40f2-a5b2-ef764b8314ee","question":"Email Address","placeholder":"Enter Your Email Id","value":"","required":false,"min":"","max":"","textareaHeight":"4rem"},{"type":"Dropdown","uniqID":"3825d8ee-b526-4f4b-ae95-19ef6e433e67","question":"Reason for Contact","options":[{"value":"General Enquiry","id":1},{"value":"Feedback","id":2},{"value":"Collaboration","id":3},{"value":"Technical Issue","id":4}],"value":"","required":false,"textareaHeight":"4rem"},{"type":"ShortAnswer","uniqID":"39f285c2-f05a-4bbd-93db-a7aa7fc4bbdc","question":"Subject","placeholder":"Enter Your Subject","value":"","required":false,"min":"","max":"","textareaHeight":"4rem"},{"type":"LongAnswer","uniqID":"0d078eb4-49ea-4cb0-bfd0-ace1e75142e0","question":"Message","placeholder":"Enter Your Message","value":"","required":false,"min":"","max":"","textareaHeight":"4rem"}],"submit":"Submit","date":"Wed Jan 22 2025 17:36:07 GMT+0530 (India Standard Time)"}',1,1,1,'current-time',0,'uu_id','/public/img/CAT-2.svg','Contact us with your questions or feedback, and we’ll respond promptly! Let me know if you need further adjustments!','Default Channel','tid','current-time',ch_id,'cta/connectus.png')
-- INSERT INTO tbl_blocks(title,block_description, block_content, cover_image, is_active , prime , created_by, created_on,tenant_id,channel_id,channel_slugname)VALUES('Success Stories','spurtcms','<section class="card-16 text-start "><div class=" mb-[4.375rem]  max-[37.5rem]:mb-[1.875rem]"><h1 class=" max-[26.5625rem]:mb-[.5rem] max-[26.5625rem]:text-[.875rem]  max-[26.5625rem]:leading-normal text-[2.25rem] font-semibold text-[#120B14] mb-[1rem] leading-10 text-center max-[37.5rem]:text-[1.5rem]">Success cases</h1><p class="max-[26.5625rem]:text-[.75rem]   max-[26.5625rem]:leading-normal text-[#151618CC] text-[1.125rem] font-normal leading-5  text-center  max-[37.5rem]:text-[1rem]">We believe spurtCMS should be accessibleto all companies, no matter the size.</p></div><div class="bg-[#ffffff] grid shadow-[0rem_0rem_1.25rem_.1875rem_rgba(36,36,36,0.04)]  grid-cols-[60%_40%] rounded-[.75rem] overflow-hidden"><div class="space-y-[1rem] p-[3.375rem] flex flex-col justify-between h-full max-[80rem]:p-[1.5rem] max-[26.5625rem]:p-[1rem]"><div class="max-[62rem]:mb-[1.125rem]"><h3 class="text-[1.5rem] text-[#120B14C4] font-semibold mb-[2.875rem] max-[62rem]:mb-[1.125rem] max-[62rem]:text-[1.25rem]  max-[37.5rem]:text-[.75rem] max-[37.5rem]:leading-normal max-[26.5625rem]:mb-[.625rem]">Svenn</h3><div><h2 class="text-[2.5rem] font-semibold text-[#120B14] mb-[1.125rem] leading-10 max-[62rem]:text-[1.75rem] max-[62rem]:leading-8  max-[37.5rem]:leading-normal max-[37.5rem]:mb-[.625rem] max-[37.5rem]:text-[.875rem] max-[26.5625rem]:text-[.75rem] max-[26.5625rem]:mb-[.3125rem]">Reducing paperwork for    construction companies</h2><p class="text-[1.125rem] font-medium text-[#1516188F]  max-[37.5rem]:text-[.75rem] max-[37.5rem]:leading-normal max-[26.5625rem]:text-[.625rem]">Construction Norway </p></div></div><a href="#" class="flex items-center space-x-1"><span class="text-[1.125rem] font-semibold text-[#120B14] max-[37.5rem]:text-[.625rem]">View case</span><img src="/public/img/viewCase.svg" alt="view case" class="max-[37.5rem]:w-[.75rem]"></a></div><div><img src="/public/img/sucessCase-img.svg" alt="sucessCase" class="w-full object-cover h-full "></div></section>','/blocks/layout19.jpg',1,0,u_id,'current-time','tid','bch_id','Blog'),('Client Portfolio','spurtcms','<section class="card-17"><h1 class="text-[2.25rem] font-semibold text-[#120B14] mb-[5.25rem] max-[80rem]:mb-[2.625rem] max-[37.5rem]:mb-[1.5rem] leading-10 text-center max-[26.5625rem]:leading-normal max-[26.5625rem]:text-[.875rem] max-[37.5rem]:text-[1.5rem]">Our Clients</h1><div class="grid grid-cols-6 items-center gap-[2rem] max-[37.5rem]:gap-[1rem]"><div><img src="/public/img/spindy.svg" alt="spindy"></div><div><img src="/public/img/gofamer.svg" alt="spindy"></div><div><img src="/public/img/momice.svg" alt="spindy"></div><div><img src="/public/img/proploy.svg" alt="spindy"></div><div><img src="/public/img/svenn.svg" alt="spindy"></div><div><img src="/public/img/bidly.svg" alt="spindy"></div><div><img src="/public/img/thesio.svg" alt="spindy"></div><div><img src="/public/img/bomvia.svg" alt="spindy"></div><div><img src="/public/img/cutis.svg" alt="spindy"></div><div><img src="/public/img/firstHome.svg" alt="spindy"></div><div><img src="/public/img/airthings.svg" alt="spindy"></div><div><img src="/public/img/hyko.svg" alt="spindy"></div><div><img src="/public/img/biderator.svg" alt="spindy"></div><div><img src="/public/img/tradesherpa.svg" alt="spindy"></div><div><img src="/public/img/sensor.svg" alt="spindy"></div><div><img src="/public/img/taskize.svg" alt="spindy"></div><div><img src="/public/img/ticketless.svg" alt="spindy"></div><div><img src="/public/img/nidarholm.svg" alt="spindy"></div><div><img src="/public/img/thesio.svg" alt="spindy"></div><div><img src="/public/img/bomvia.svg" alt="spindy"></div><div><img src="/public/img/cutis.svg" alt="spindy"></div><div><img src="/public/img/firstHome.svg" alt="spindy"></div><div><img src="/public/img/airthings.svg" alt="spindy"></div><div><img src="/public/img/hyko.svg" alt="spindy"></div></div></section>','/blocks/layout20.jpg',1,0,u_id,'current-time','tid','jch_id','Jobs'),('industry expertise','spurtcms','<section class="card-18 text-start"><h1 class="text-[2.25rem] font-semibold text-[#120B14] mb-[4.3125rem] max-[80rem]:mb-[2.625rem] max-[37.5rem]:mb-[1.5rem] leading-10 text-center max-[26.5625rem]:leading-normal max-[26.5625rem]:text-[.875rem] max-[37.5rem]:text-[1.5rem]">Industry experise</h1><div class="bg-[url(''/public/img/real-estate.svg'')] bg-cover rounded-[.75rem] overflow-hidden"><div class="bg-[linear-gradient(268.55deg,_rgba(37,37,37,0)_0.51%,_#020202_127.16%)] p-[2.1875rem] max-[48rem]:p-[1rem] max-[48rem]:min-h-[auto] min-h-[22.8125rem] flex flex-col h-full"><h1 class="max-[37.5rem]:text-[.875rem] text-[2.125rem] max-[48rem]:mb-[1.125rem] max-[48rem]:text-[1.5rem] font-bold text-[#FFFFFF] mb-[1.875rem]">Real Estate</h1><ul class="ps-[2rem] mb-[3.375rem] max-[48rem]:mb-[1.25rem]"><li class="max-[37.5rem]:text-[.75rem] max-[48rem]:text-[1rem] text-[1.375rem] text-[#FFFFFF] font-medium mb-[.625rem] leading-normal list-disc">Search and recommendation engine</li><li class="max-[37.5rem]:text-[.75rem] max-[48rem]:text-[1rem] text-[1.375rem] text-[#FFFFFF] font-medium mb-[.625rem] leading-normal list-disc">Real estate valuation software</li><li class="max-[37.5rem]:text-[.75rem] max-[48rem]:text-[1rem] text-[1.375rem] text-[#FFFFFF] font-medium mb-[.625rem] leading-normal list-disc">Portals for owner and tenants</li><li class="max-[37.5rem]:text-[.75rem] max-[48rem]:text-[1rem] text-[1.375rem] text-[#FFFFFF] font-medium leading-normal list-disc">Apps for real agent management</li></ul><a href="#" class="max-[37.5rem]:text-[.625rem] max-[48rem]:text-[.875rem] mt-auto flex items-center space-x-[.25rem] text-[1.125rem] font-bold text-[#FFFFFF]"><span>Learn More</span><img src="/public/img/learn-more.svg" alt="learn more" class="max-[37.5rem]:w-[.625rem]"></a></div></div></section>','/blocks/layout21.jpg',1,0,u_id,'current-time','tid','jch_id','Jobs')




-- INSERT INTO tbl_integrations (gateway_name,gateway_desc,cover_image,client_id,client_secret,is_active,is_deleted,created_by,created_on,tenant_id)VALUES('PayPal','A trusted global payment system for sending, receiving, and managing money.','/public/img/paypal.svg','pk_test_51QzENzEVUSWnoB21KWRGzk8hUY1nxJpDBb9C0Ev4imAKqRp1kcm7zHEMljwi75HKIiyTgXvQbwECRv4ml74ufmrZ00CPHdbq8v','sk_test_51QzENzEVUSWnoB21CqdQibHzjqwFDN0y2VIgmPb6VCIZHYLyQlJXd9Yp2ozUhyvKg4SiFLSqoCBMVcr3kVCo1v0d00MkEqkAWM',1,0,1,'current-time','tid')
-- INSERT INTO tbl_integrations (gateway_name,gateway_desc,cover_image,client_id,client_secret,is_active,is_deleted,created_by,created_on,tenant_id)VALUES('Stripe','A powerful financial platform enabling online and in-person transactions.','/public/img/stripe.svg','pk_test_51QzENzEVUSWnoB21KWRGzk8hUY1nxJpDBb9C0Ev4imAKqRp1kcm7zHEMljwi75HKIiyTgXvQbwECRv4ml74ufmrZ00CPHdbq8v','sk_test_51QzENzEVUSWnoB21CqdQibHzjqwFDN0y2VIgmPb6VCIZHYLyQlJXd9Yp2ozUhyvKg4SiFLSqoCBMVcr3kVCo1v0d00MkEqkAWM',1,0,1,'current-time','tid')


Insert Into tbl_mstr_membergrouplevels(group_name,description,slug,created_on,created_by,is_deleted,is_active,tenant_id) VALUES ('Default Group','Default Membership Group','default_group','current-time', u_id,0,1,'tid')

Insert Into tbl_mstr_membershiplevels(subscription_name,description,membershiplevel_details,initial_payment,created_on,created_by,is_deleted,is_active,tenant_id) VALUES ('Explorer','Begin your journey with the essentials','Your Membership Information You have chosen the Explorer Membership.',100,'current-time', u_id,0,1,'tid')

Insert Into tbl_courses (title,description,category_id,status,tenant_id,created_on,created_by,is_deleted,uuid) VALUES ('Default Course','Default Course',co_id,0,'tid','current-time', u_id,0,'duid')

Insert Into tbl_course_settings (course_id,certificate,comments,offer,tenant_id,created_on,created_by,is_deleted) VALUES (cos_id,0,0,'Free','tid','current-time', u_id,0)



-- INSERT INTO tbl_menus(name,description,slug_name, status,parent_id,tenant_id,created_on,created_by,is_deleted)VALUES('Header','The primary menu at the top of a webpage, guiding users to major site sections','header',1,0,'tid','current-time', u_id,0)

