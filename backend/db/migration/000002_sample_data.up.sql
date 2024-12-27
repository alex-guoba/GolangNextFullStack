INSERT INTO `blog_tag` (`id`, `name`, `created_by`, `modified_by`, `ts_create`, `ts_update`, `is_del`, `state`) VALUES
(1, 'Technology', 'Admin', 'Admin', NOW(), NOW(), 0, 1),
(2, 'Lifestyle', 'Admin', 'Admin', NOW(), NOW(), 0, 1),
(3, 'Health', 'Admin', 'Admin', NOW(), NOW(), 0, 1),
(4, 'Programming', 'Admin', 'Admin', NOW(), NOW(), 0, 1),
(5, 'Fitness', 'Admin', 'Admin', NOW(), NOW(), 0, 1),
(6, 'Travel', 'Admin', 'Admin', NOW(), NOW(), 0, 1);

INSERT INTO `blog_article` (`id`, `title`, `desc`, `cover_image_url`, `content`, `created_by`, `modified_by`, `ts_create`, `ts_update`, `is_del`, `state`) VALUES
(1, 'Introduction to AI', 'An overview of artificial intelligence.', 'https://example.com/image1.jpg', 'Content about AI...', 'Author1', 'Author1', NOW(), NOW(), 0, 1),
(2, 'Staying Healthy in the Digital Age', 'Tips for maintaining health in a tech-driven world.', 'https://example.com/image2.jpg', 'Content about health...', 'Author2', 'Author2', NOW(), NOW(), 0, 1),
(3, 'Exploring New Destinations', 'A guide to discovering new travel spots.', 'https://example.com/image3.jpg', 'Content about travel...', 'Author3', 'Author3', NOW(), NOW(), 0, 1),
(4, 'Advanced JavaScript Techniques', 'Learn advanced JavaScript programming.', 'https://example.com/image4.jpg', 'Content about JavaScript...', 'Author1', 'Author1', NOW(), NOW(), 0, 1),
(5, 'Building Muscle at Home', 'Workout routines to build muscle without a gym.', 'https://example.com/image5.jpg', 'Content about fitness...', 'Author2', 'Author2', NOW(), NOW(), 0, 1),
(6, 'The Future of Work', 'Predictions about the future workplace.', 'https://example.com/image6.jpg', 'Content about future work...', 'Author3', 'Author3', NOW(), NOW(), 0, 1),
(7, 'The Benefits of Meditation', 'How meditation can improve your life.', 'https://example.com/image7.jpg', 'Content about meditation...', 'Author1', 'Author1', NOW(), NOW(), 0, 1),
(8, 'The Importance of Sleep', 'Why getting enough sleep is crucial.', 'https://example.com/image8.jpg', 'Content about sleep...', 'Author2', 'Author2', NOW(), NOW(), 0, 1),
(9, 'Cooking Healthy Meals', 'Recipes and tips for healthy eating.', 'https://example.com/image9.jpg', 'Content about healthy cooking...', 'Author3', 'Author3', NOW(), NOW(), 0, 1),
(10, 'Starting a Side Hustle', 'Advice for starting a part-time business.', 'https://example.com/image10.jpg', 'Content about side hustle...', 'Author1', 'Author1', NOW(), NOW(), 0, 1),
(11, 'Improving Your Memory', 'Techniques to enhance memory and recall.', 'https://example.com/image11.jpg', 'Content about memory...', 'Author2', 'Author2', NOW(), NOW(), 0, 1),
(12, 'Understanding Quantum Computing', 'An introduction to quantum computing.', 'https://example.com/image12.jpg', 'Content about quantum computing...', 'Author3', 'Author3', NOW(), NOW(), 0, 1);


INSERT INTO `blog_article_tag` (`id`, `article_id`, `tag_id`, `created_by`, `modified_by`, `ts_create`, `ts_update`, `is_del`) VALUES
(1, 1, 4, 'Admin', 'Admin', NOW(), NOW(), 0),
(2, 1, 1, 'Admin', 'Admin', NOW(), NOW(), 0),
(3, 2, 2, 'Admin', 'Admin', NOW(), NOW(), 0),
(4, 2, 5, 'Admin', 'Admin', NOW(), NOW(), 0),
(5, 3, 6, 'Admin', 'Admin', NOW(), NOW(), 0),
(6, 4, 4, 'Admin', 'Admin', NOW(), NOW(), 0),
(7, 5, 5, 'Admin', 'Admin', NOW(), NOW(), 0),
(8, 6, 1, 'Admin', 'Admin', NOW(), NOW(), 0),
(9, 7, 2, 'Admin', 'Admin', NOW(), NOW(), 0),
(10, 8, 5, 'Admin', 'Admin', NOW(), NOW(), 0),
(11, 9, 2, 'Admin', 'Admin', NOW(), NOW(), 0),
(12, 10, 3, 'Admin', 'Admin', NOW(), NOW(), 0),
(13, 11, 5, 'Admin', 'Admin', NOW(), NOW(), 0),
(14, 12, 1, 'Admin', 'Admin', NOW(), NOW(), 0),
(15, 12, 4, 'Admin', 'Admin', NOW(), NOW(), 0);