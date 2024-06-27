-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jun 27, 2024 at 11:17 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `farah`
--

-- --------------------------------------------------------

--
-- Table structure for table `ry_abilities`
--

CREATE TABLE `ry_abilities` (
  `ryabilit_id` int(11) NOT NULL,
  `ryabilit_name` varchar(255) DEFAULT NULL,
  `ryabilit_icon` varchar(255) DEFAULT NULL,
  `ryabilit_type` varchar(255) DEFAULT NULL,
  `ryabilit_crdate` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `ry_abilities`
--

INSERT INTO `ry_abilities` (`ryabilit_id`, `ryabilit_name`, `ryabilit_icon`, `ryabilit_type`, `ryabilit_crdate`) VALUES
(1, 'Garden', 'garden', 'main', '2023-09-30 12:02:49'),
(2, 'Terrace', 'terrace', 'main', '2023-09-30 12:02:49'),
(3, 'Garage', 'garage', 'main', '2023-09-30 12:02:49'),
(4, 'Elevator', 'elevator', 'main', '2023-09-30 12:02:49'),
(5, 'Sea views', 'sea-view', 'main', '2023-09-30 12:02:49'),
(6, 'Mountains view', 'mountains-view', 'main', '2023-09-30 12:02:49'),
(7, 'Pool', 'pool', 'main', '2023-09-30 12:02:49'),
(8, 'Concierge', 'concierge', 'main', '2023-09-30 12:02:49'),
(9, 'Storage room', 'storage-room', 'main', '2023-09-30 12:02:49'),
(10, 'Salon', 'salon', 'inner', '2023-09-30 12:02:49'),
(11, 'Satellite dish', 'satellite-dish', 'inner', '2023-09-30 12:02:49'),
(12, 'Fireplace', 'fireplace', 'inner', '2023-09-30 12:02:49'),
(13, 'Air conditioning', 'air-conditioner', 'inner', '2023-09-30 12:02:49'),
(14, 'Heating', 'heater', 'inner', '2023-09-30 12:02:49'),
(15, 'Security system', 'cctv-camera', 'inner', '2023-09-30 12:02:49'),
(16, 'Double glazing', 'double-glazing', 'inner', '2023-09-30 12:02:49'),
(17, 'Reinforced Door', 'smart-door', 'inner', '2023-09-30 12:02:49'),
(18, 'Equipped kitchen', 'kitchen', 'additional', '2023-09-30 12:02:49'),
(19, 'Fridge', 'fridge', 'additional', '2023-09-30 12:02:49'),
(20, 'Oven', 'oven', 'additional', '2023-09-30 12:02:49'),
(21, 'TV', 'tv', 'additional', '2023-09-30 12:02:49'),
(22, 'Washing machine', 'washing-machine', 'additional', '2023-09-30 12:02:49'),
(23, 'Microwave', 'microwave', 'additional', '2023-09-30 12:02:49'),
(24, 'Internet', 'internet', 'additional', '2023-09-30 12:02:49');

-- --------------------------------------------------------

--
-- Table structure for table `ry_annonce`
--

CREATE TABLE `ry_annonce` (
  `royaannonse_userid` int(11) DEFAULT NULL,
  `royaannonse_region` varchar(255) DEFAULT NULL,
  `royaannonse_city` varchar(255) DEFAULT NULL,
  `royaannonse_transaction` varchar(255) DEFAULT NULL,
  `royaannonse_propertyType` varchar(255) DEFAULT NULL,
  `royaannonse_status` int(11) DEFAULT 0,
  `royaannonse_address` varchar(255) DEFAULT NULL,
  `royaannonse_quartier` varchar(255) DEFAULT NULL,
  `royaannonse_area` varchar(255) DEFAULT NULL,
  `royaannonse_price` varchar(255) DEFAULT NULL,
  `royaannonse_age` varchar(255) DEFAULT NULL,
  `royaannonse_floorType` varchar(255) DEFAULT NULL,
  `royaannonse_floor` varchar(255) DEFAULT NULL,
  `royaannonse_apartment` varchar(255) DEFAULT NULL,
  `royaannonse_bedrooms` int(11) DEFAULT NULL,
  `royaannonse_bathrooms` int(11) DEFAULT NULL,
  `royaannonse_kitchens` int(11) DEFAULT NULL,
  `royaannonse_title` varchar(255) DEFAULT NULL,
  `royaannonse_description` varchar(1024) DEFAULT NULL,
  `royaannonse_phone1` varchar(255) DEFAULT NULL,
  `royaannonse_phone2` varchar(255) DEFAULT NULL,
  `royaannonse_phone3` varchar(255) DEFAULT NULL,
  `royaannonse_validated` int(11) DEFAULT 0,
  `royaannonse_is_delete` int(11) DEFAULT 0,
  `royaannonse_crdate` timestamp NOT NULL DEFAULT current_timestamp(),
  `royaannonse_deleted_on` datetime DEFAULT NULL,
  `royaannonse_deleted_by` int(11) DEFAULT NULL,
  `royaannonse_id` int(11) NOT NULL,
  `royaannonse_category` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_annonce`
--

INSERT INTO `ry_annonce` (`royaannonse_userid`, `royaannonse_region`, `royaannonse_city`, `royaannonse_transaction`, `royaannonse_propertyType`, `royaannonse_status`, `royaannonse_address`, `royaannonse_quartier`, `royaannonse_area`, `royaannonse_price`, `royaannonse_age`, `royaannonse_floorType`, `royaannonse_floor`, `royaannonse_apartment`, `royaannonse_bedrooms`, `royaannonse_bathrooms`, `royaannonse_kitchens`, `royaannonse_title`, `royaannonse_description`, `royaannonse_phone1`, `royaannonse_phone2`, `royaannonse_phone3`, `royaannonse_validated`, `royaannonse_is_delete`, `royaannonse_crdate`, `royaannonse_deleted_on`, `royaannonse_deleted_by`, `royaannonse_id`, `royaannonse_category`) VALUES
(1, NULL, 'Agadir', 'Vente', 'Bon état', 0, 'ايت ملول ازرو', '', '200', '8000000', 'Moins de 1 ans', NULL, '', '', 4, 4, 4, 'فيلا رائعة للبيع', 'فيلا رائعة للبيع . 4 غرف ممتازة. بيئة هادئة مع إطلالة على الجبل ،التدفئة المركزية.', '', '', '', 0, 0, '2023-10-05 19:45:03', NULL, NULL, 1, 5),
(1, NULL, 'Marrakech', 'Vente', 'Bon état', 0, 'Marrakech, Ourika Road', NULL, '200', '300000', '1 ans', NULL, NULL, NULL, 4, 2, 1, 'House for sale in Route de Ourika', 'Very nice house for sale in Route de l&#39;Ourika. 4 beautiful rooms. caretaker available, air conditioning system.', NULL, NULL, NULL, 0, 0, '2024-06-22 23:24:56', NULL, NULL, 3, 2),
(1, NULL, 'Marrakech', 'Vente', 'Bon état', 0, 'Mrrakech ....', NULL, '400', '2000000', '1 ans', NULL, NULL, NULL, 4, 2, 2, 'Single storey villa km8', 'We offer you this sublime single storey villa less than 7 minutes from the city center. Built on a plot of 2000m2, offering 600m2 covered space and 4 suites + office, gym, hammam and sauna. Oriented south south west, this bright, chic and functional villa without any opposite will seduce you with its volumes and its sobriety. Beautiful finishes, reversible air conditioning, double glazing, electrolysis infinity pool. Contact us for further information or to visit this exceptional property. Compulsory visit voucher. Agency fees 2.5% excluding tax. Price upon request.', NULL, NULL, NULL, 0, 0, '2024-06-24 18:57:53', NULL, NULL, 4, 5),
(1, NULL, 'Marrakech', 'Vente', 'Bon état', 0, 'Marrakech', NULL, '60', '2420000', '1 ans', NULL, NULL, NULL, 3, 1, 1, 'Beautiful apartment for sale', 'Buy your apartment. Price 2,420,000 DH. 2 rooms, 1 bathroom, surface area 60 m². 1 bedroom. New. Enjoy a beautiful terrace. Also has an elevator.\r\n\r\nCome and discover this apartment for sale in Hivernage, Marrakech. Big pool.', NULL, NULL, NULL, 0, 0, '2024-06-24 19:06:42', NULL, NULL, 5, 1);

-- --------------------------------------------------------

--
-- Table structure for table `ry_annonce_abilities_avariable`
--

CREATE TABLE `ry_annonce_abilities_avariable` (
  `ryannoabilitvarb_id` int(11) NOT NULL,
  `ryannoabilitvarb_annonce_id` int(11) DEFAULT NULL,
  `ryannoabilitvarb_abilitie_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `ry_annonce_abilities_avariable`
--

INSERT INTO `ry_annonce_abilities_avariable` (`ryannoabilitvarb_id`, `ryannoabilitvarb_annonce_id`, `ryannoabilitvarb_abilitie_id`) VALUES
(1, 1, 1),
(2, 1, 2),
(3, 1, 3),
(4, 1, 4),
(5, 1, 5),
(6, 1, 6),
(7, 1, 7),
(8, 1, 8),
(9, 1, 9),
(10, 1, 10),
(11, 1, 11),
(12, 1, 12),
(13, 1, 13),
(14, 1, 14),
(15, 1, 15),
(16, 1, 16),
(17, 1, 17),
(18, 1, 18),
(19, 1, 19),
(20, 1, 20),
(21, 1, 21),
(22, 1, 22),
(23, 1, 23),
(24, 1, 24),
(25, 3, 17),
(26, 3, 18),
(27, 3, 19),
(28, 3, 21),
(29, 3, 22),
(30, 3, 23),
(31, 3, 24),
(32, 3, 17),
(33, 3, 13),
(34, 3, 17),
(35, 3, 18),
(36, 3, 19),
(37, 3, 21),
(38, 3, 22),
(39, 3, 23),
(40, 3, 24),
(41, 0, 1),
(42, 0, 2),
(43, 0, 3),
(44, 0, 5),
(45, 0, 12),
(46, 0, 18),
(47, 0, 19),
(48, 0, 20),
(49, 0, 22),
(50, 0, 23),
(51, 0, 24),
(52, 0, 1),
(53, 0, 2),
(54, 0, 3),
(55, 0, 5),
(56, 0, 12),
(57, 0, 18),
(58, 0, 19),
(59, 0, 20),
(60, 0, 22),
(61, 0, 23),
(62, 0, 24),
(63, 0, 1),
(64, 0, 2),
(65, 0, 3),
(66, 0, 5),
(67, 0, 12),
(68, 0, 18),
(69, 0, 19),
(70, 0, 20),
(71, 0, 22),
(72, 0, 23),
(73, 0, 24),
(74, 4, 1),
(75, 4, 2),
(76, 4, 3),
(77, 4, 5),
(78, 4, 12),
(79, 4, 18),
(80, 4, 19),
(81, 4, 20),
(82, 4, 22),
(83, 4, 23),
(84, 4, 24),
(85, 5, 2),
(86, 5, 3),
(87, 5, 4),
(88, 5, 5),
(89, 5, 8),
(90, 5, 24);

-- --------------------------------------------------------

--
-- Table structure for table `ry_annonce_images_avariable`
--

CREATE TABLE `ry_annonce_images_avariable` (
  `royaannimgava_id` int(11) NOT NULL,
  `royaannimgava_annonce_id` int(11) NOT NULL,
  `royaannimgava_image_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_annonce_images_avariable`
--

INSERT INTO `ry_annonce_images_avariable` (`royaannimgava_id`, `royaannimgava_annonce_id`, `royaannimgava_image_id`) VALUES
(1, 1, 1),
(2, 1, 2),
(3, 1, 3),
(4, 1, 4),
(5, 3, 7),
(6, 3, 8),
(7, 3, 9),
(8, 0, 10),
(9, 0, 11),
(10, 0, 12),
(11, 4, 13),
(12, 4, 14),
(13, 4, 15),
(14, 5, 16),
(15, 5, 17),
(16, 5, 18);

-- --------------------------------------------------------

--
-- Table structure for table `ry_annonce_message`
--

CREATE TABLE `ry_annonce_message` (
  `ryannmess_id` int(11) NOT NULL,
  `ryannmess_user_id` int(11) NOT NULL,
  `ryannmess_admin_id` int(11) NOT NULL,
  `ryannmess_annonce_id` int(11) NOT NULL,
  `ryannmess_message` varchar(255) DEFAULT NULL,
  `ryannmess_crdate` timestamp NULL DEFAULT current_timestamp(),
  `ryannmess_create_by` int(11) NOT NULL,
  `ryannmess_status` int(11) NOT NULL DEFAULT 0,
  `ryannmess_seemessage` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `ry_annonce_message`
--

INSERT INTO `ry_annonce_message` (`ryannmess_id`, `ryannmess_user_id`, `ryannmess_admin_id`, `ryannmess_annonce_id`, `ryannmess_message`, `ryannmess_crdate`, `ryannmess_create_by`, `ryannmess_status`, `ryannmess_seemessage`) VALUES
(1, 3, 1, 1, 'ouknik you are a bad student', '2023-10-01 16:27:31', 3, 1, '2024-06-24 19:33:55'),
(2, 6, 1, 1, 'ouknik you are a bad student', '2023-10-01 16:28:23', 6, 1, '2024-06-24 19:33:56'),
(3, 6, 1, 1, 'ouknik you are a bad student', '2023-10-01 16:40:02', 6, 1, '2024-06-24 19:33:56'),
(4, 6, 1, 1, 'ouknik you are a bad student', '2023-10-01 16:42:49', 6, 1, '2024-06-24 19:33:56'),
(5, 6, 1, 1, 'ouknik you are a bad student', '2023-10-01 16:44:55', 6, 1, '2024-06-24 19:33:56'),
(6, 6, 1, 1, 'ouknik you are a bad student', '2023-10-01 16:45:43', 6, 1, '2024-06-24 19:33:56'),
(7, 6, 1, 1, 'ouknik', '2023-10-02 10:33:29', 6, 1, '2024-06-24 19:33:56'),
(8, 6, 1, 1, 'ouknik', '2023-10-02 10:40:22', 6, 1, '2024-06-24 19:33:56'),
(9, 6, 1, 1, 'thats now is good', '2023-10-02 10:52:33', 6, 1, '2024-06-24 19:33:56'),
(10, 3, 1, 1, 'thats is bad from you', '2023-10-02 10:53:33', 3, 1, '2024-06-24 19:33:55'),
(11, 3, 1, 1, 'thats is bad from you', '2023-10-02 12:25:38', 3, 1, '2024-06-24 19:33:55'),
(12, 3, 1, 1, 'yes I now thinks bro', '2023-10-02 12:26:10', 1, 1, '2023-10-03 13:48:44'),
(13, 3, 1, 1, 'fack you', '2023-10-02 12:31:46', 1, 1, '2023-10-03 13:48:45'),
(14, 3, 1, 1, 'I know this but', '2023-10-02 12:34:43', 1, 1, '2023-10-03 13:48:44'),
(15, 3, 1, 1, 'ok ok', '2023-10-02 12:36:02', 1, 1, '2023-10-03 13:48:43'),
(16, 3, 1, 1, 'yes', '2023-10-02 12:37:16', 1, 1, '2023-10-03 13:48:43'),
(17, 3, 1, 1, 'no thanks', '2023-10-02 12:43:56', 1, 1, '2023-10-03 13:48:43'),
(18, 3, 1, 1, 'ok ok', '2023-10-02 12:45:50', 1, 1, '2023-10-03 13:48:43'),
(19, 6, 1, 1, 'okay thanks', '2023-10-02 12:51:04', 1, 0, NULL),
(20, 1, 3, 1, 'thats is bad from you', '2023-10-02 13:00:57', 1, 1, '2023-10-02 15:51:11'),
(21, 3, 1, 1, 'thats', '2023-10-02 13:11:39', 3, 1, '2024-06-24 19:33:55'),
(22, 1, 3, 1, 'thats', '2023-10-02 13:12:19', 1, 1, '2023-10-02 15:51:12'),
(23, 3, 1, 1, 'thats', '2023-10-02 13:13:00', 1, 1, '2023-10-03 13:48:45'),
(24, 6, 1, 1, 'yes', '2023-10-02 13:15:12', 1, 0, NULL),
(25, 3, 1, 1, 'ok ok', '2023-10-02 13:17:57', 1, 1, '2023-10-03 13:48:45'),
(26, 3, 1, 1, 'ok ok', '2023-10-02 13:18:50', 1, 1, '2023-10-03 13:48:44'),
(27, 3, 1, 1, 'test', '2023-10-02 14:07:38', 3, 1, '2024-06-24 19:33:55'),
(28, 1, 3, 1, 'ok', '2023-10-02 14:39:47', 3, 0, NULL),
(29, 1, 3, 1, 'Fack you', '2023-10-02 14:42:14', 3, 0, NULL),
(30, 3, 1, 1, 'ok thats good from you bro', '2023-10-02 14:52:21', 3, 1, '2024-06-24 19:33:55'),
(31, 3, 1, 1, 'I need info for thi annonce', '2023-10-02 14:56:32', 3, 1, '2024-06-24 19:33:55'),
(32, 3, 1, 43, 'sal', '2023-10-02 22:00:09', 3, 1, '2024-06-24 19:33:55'),
(33, 3, 1, 43, 'sal', '2023-10-02 22:01:16', 3, 1, '2024-06-24 19:33:55'),
(34, 3, 1, 45, 'yes', '2023-10-03 10:23:50', 3, 1, '2024-06-24 19:33:55'),
(35, 3, 1, 1, 'test', '2023-10-03 11:14:04', 3, 1, '2024-06-24 19:33:55'),
(36, 3, 1, 1, 'yes maybe', '2023-10-03 11:19:49', 1, 1, '2023-10-03 13:48:45'),
(37, 3, 1, 1, 'yes  of cource', '2023-10-03 11:42:36', 3, 1, '2024-06-24 19:33:55'),
(38, 3, 1, 1, 'hi', '2023-10-03 12:48:54', 3, 1, '2024-06-24 19:33:55'),
(39, 3, 1, 1, 'ho', '2023-10-04 10:46:35', 1, 0, NULL),
(40, 6, 1, 1, 'ouknik', '2023-10-05 12:33:27', 6, 1, '2024-06-24 19:33:56'),
(41, 6, 1, 1, 'yes ok sir thats is good', '2023-10-06 18:43:46', 1, 0, NULL),
(42, 3, 1, 1, 'ouknik is good men', '2023-10-06 19:06:41', 3, 1, '2024-06-24 19:33:55'),
(43, 3, 1, 1, 'gdhffg', '2023-10-06 19:07:17', 1, 0, NULL),
(44, 3, 1, 1, 'ouknik is good men', '2023-10-06 19:22:13', 3, 1, '2024-06-24 19:33:55'),
(45, 4, 1, 1, 'ouknik is good men', '2023-10-06 19:22:56', 4, 1, '2024-06-24 19:33:52'),
(46, 2, 1, 1, 'ouknik is good men', '2023-10-06 19:23:59', 2, 1, '2024-06-24 19:33:38');

-- --------------------------------------------------------

--
-- Table structure for table `ry_annose_images`
--

CREATE TABLE `ry_annose_images` (
  `royaimg_id` int(11) NOT NULL,
  `royaimg_image` varchar(10000) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_annose_images`
--

INSERT INTO `ry_annose_images` (`royaimg_id`, `royaimg_image`) VALUES
(1, '/p7Ng_entree résidence villas park avenue marrakech _49874498.jpeg'),
(2, '/0rYP_type 2 residence park avenue marrakech_50622736.jpeg'),
(3, '/x7IU_WhatsApp Image 2023-06-09 at 16.38.37_50622737.jpeg'),
(4, '/gVMr_villa type 2 residence park avenue marrakech_50622734.jpeg'),
(5, 'restntimg_dOMQ_2023-02-15T195109-504Z-Bureau-0-chambres-GrandCasablanca-Settat-SidiBernoussi1676494331412-613edb83-0a7f-4ccf-b9d3-38e6960f5c0b-annonce-agenz.jpeg'),
(6, 'restntimg_s8y1_WhatsApp Image 2024-01-25 at 00.03.35(4).jpeg'),
(7, 'restntimg_cMW9_WhatsApp Image 2024-01-25 at 00.03.35(4).jpeg'),
(8, 'restntimg_Gn3p_10097513237.jpg'),
(9, 'restntimg_QFmD_WhatsApp Image 2024-01-25 at 00.03.33(1).jpeg'),
(10, 'restntimg_Zi4b_Villa-Aaliyah-Marrakesh-Exterior.JPEG'),
(11, 'restntimg_Z1LW_Photo-24-03-2014-19-46-14-copie.jpg'),
(12, 'restntimg_g5tW_villa-marrakech-belya-01.jpg'),
(13, 'restntimg_I48N_villa-marrakech-belya-01.jpg'),
(14, 'restntimg_CQcU_Photo-24-03-2014-19-46-14-copie.jpg'),
(15, 'restntimg_cimf_Villa-Aaliyah-Marrakesh-Exterior.JPEG'),
(16, 'restntimg_7zON_176245280.jpg'),
(17, 'restntimg_YIeT_176394968.jpg'),
(18, 'restntimg_8tdA_vente-appartement-marrakech-centre-ville1.jpeg');

-- --------------------------------------------------------

--
-- Table structure for table `ry_api_otts`
--

CREATE TABLE `ry_api_otts` (
  `royaapiot_id` varchar(64) NOT NULL,
  `royaapiot_api_ott` varchar(8) DEFAULT NULL,
  `royaapiot_device_id` bigint(20) DEFAULT NULL,
  `royaapiot_ip` varchar(15) DEFAULT NULL,
  `royaapiot_used_on` varchar(255) DEFAULT NULL,
  `royaapiot_crdate` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_api_otts`
--

INSERT INTO `ry_api_otts` (`royaapiot_id`, `royaapiot_api_ott`, `royaapiot_device_id`, `royaapiot_ip`, `royaapiot_used_on`, `royaapiot_crdate`) VALUES
('1S5w8iZzwC88RDFPVPQb5gul3BxMKkXOXsBPs7JY8rOQoC4qQV8s862rBBqaJITX', 'nnPKhcFX', 8, NULL, NULL, '2024-06-24 18:23:37'),
('3GyQ5OLtTv3rnHJR6Q4ri5CBGYfdV04hP1zzcd4lfydsASidY12D5NM2CghfIJPQ', '8s77Ujz1', 1, NULL, NULL, '2024-06-24 18:28:49'),
('3hxTt2fu9PA5y08fYB9eLzzd9TwoctVlWwuihBTbdx8bAuzqjv5ND6N2slfobgKP', '6OmTJ2IB', 1, NULL, NULL, '2024-06-24 18:30:12'),
('4C7mRHHQAd9Cq4Xp4tXOFZaruaF8RhQi0DZDcD55ePnUd3fGO7mfKgENf6MKNLLe', 'y7fQOLsY', 1, NULL, NULL, '2024-06-24 18:33:53'),
('5zZRmb6d0GIomdyoB7OG2IQY7NnFdfMgwqwHT0jcogDiRYwBcROtBA0d9bMCEoRN', 'FnqsSbJ4', 1, NULL, NULL, '2024-06-25 14:33:43'),
('8PQWZqSOLXnyGwrBhp1cg3mZAqdlENCLMYI1ZJZgD3nQ4OI7MNSA9AzKTWKPGSFg', 'sWstlNN5', 1, NULL, NULL, '2024-06-24 18:33:30'),
('9Fc1EXwmFyt3hrYKvrVbB0NVdF2Qxzfi0Bgo8qXrp4jSW9LAqJbswZRo4m4Bzhud', 'CtpVW4mh', 1, NULL, NULL, '2024-06-25 14:33:51'),
('AfzYjKf8tQrduqPAjw0yhUC4QcLUBbFCmHXoEFApUp5kMum4Mp249ErilqyNdS1Q', '7bvczb6C', 1, NULL, NULL, '2024-06-25 08:21:00'),
('AIabuK8LeJTatWaTJ7kVE97gDYtREktCKzemzLyiIGQxH5yJv9aGkNDG6wx0c01S', 'V9e8BcJo', 1, NULL, NULL, '2024-06-25 14:33:16'),
('BKKXFvTPKUJsFdM9kRCaM8CW1RyZX8EaItKwySNCS0IMhfxaElZgA1J7lInVnkBx', '0JXQJUdb', 1, NULL, NULL, '2024-06-25 14:33:16'),
('BNJQy3Lw6CGHZIqcrr7SAPxzwoKUMfSWwU6okMqaDwfxqIODBIPl6yxrhWT6HdTd', 'gd5KoWYZ', 1, NULL, NULL, '2024-06-25 08:22:38'),
('C5KiWn6NIbP78ItrlKqoofI2OpLTTA5oOQfVLJdtyijyofRR07ZBurgG5oNVNwNJ', '4znZpJQB', 1, NULL, NULL, '2024-06-24 18:28:49'),
('cdMYf9sQHJc7m8Djq5qUsRmzg4AT0oN0FFGy7nnQvYINqGkGHvTE5uu0Ez3dWvhR', 'YROttkNe', 1, NULL, NULL, '2024-06-24 18:33:36'),
('e2KxWlM1vCvtJ8mFgVlaoVCf9cYNg2OqLkofC6f6PZEYsbLizB5FVzrqfCSr2taW', 'loOMLZWz', 8, NULL, NULL, '2024-06-25 08:19:07'),
('eSz5YRpPurlcXr1n8tSQDZxwWrZIDycHAqUQQRXJjsC2mMjShQdJL308NOVhuQ09', 'Z0HKVhiE', 1, NULL, NULL, '2024-06-24 18:33:55'),
('EuA92Z0tyyumiP0LviwWXYCRZUX9RuyOD2JLOQWIQLaoLkUGMzsHs9d3DnHKkAdM', 'ZqNHXBDl', 1, NULL, NULL, '2024-06-24 18:33:51'),
('fAbgVuhaMCOvCB3kDmFv6oMJob41RtaYWokUNllNdP3YL7fcpxR8M06YmUYC0Gnn', 'bQ9Skp2X', 8, NULL, NULL, '2024-06-24 18:28:45'),
('GvjecUOEIt9tr3jladvwUVOxNrnQOQtFY0cbVcYV6Ldl9lc7K1rxhzbZ2v2C6XBj', 'hsJftT28', 1, NULL, NULL, '2023-10-14 16:44:44'),
('iA9sDMu5wQ1weDuAiZXJ6pADYru6B8VLX5b7OmWqtI2BorxRyyshvoTIG03cHpKO', 'Ju4tuOco', 1, NULL, NULL, '2024-06-24 18:31:10'),
('IBWJKimICIn63OvMjhMtzU5T0ZaVbzBzxkItMwnyUHRry2MuDBxkAP3TEefBy7zU', 'cxwNuF5r', 1, NULL, NULL, '2024-06-24 18:33:51'),
('iokGLFeuOz0ygc3EKhOX850uDWVd2nOKUovIjE96WL282W2UnUBjgrBNPmFS0LRF', 'Tix5AntW', 1, NULL, NULL, '2024-06-25 08:20:34'),
('j2cec2JBGsWkKRqBnqOhymAWuMX4eeABqnvfSRM9IVRQBdPbRzxgyoA8XAwLylDq', 'sSgrEmTl', 1, NULL, NULL, '2024-06-24 18:29:22'),
('k8ACoBRSMOuAFDAXz3wPIzgS5i3TA8m9wgHXNrEhveRSUUnVXNDe7RWKXKmNXD4R', 'c2H6jElK', 1, NULL, NULL, '2024-06-24 18:33:36'),
('KrMlk0KLkhW3frhuusZPB5dA6hEM2N36ajusjrqJj31Ry9GdiEA95ijyxp9yoTlr', 'xO0pSeGt', 1, NULL, NULL, '2024-06-24 18:28:49'),
('lnRRMZzIgACXWSKW4eeFu9onJtlobCo0w0pcoKDc6uQf97cw0KMlqOwYXaa5bFrs', '1igThvEq', 1, NULL, NULL, '2024-06-24 18:23:41'),
('MxnmYm5osDM0AkavtkUwpMARGcFhAtKLNyWlx8PDCIedFPA8TkupR544BWwzMyOd', '83OMGHXc', 8, NULL, NULL, '2024-06-24 18:28:26'),
('n0PZMSZqADEbLN0cnyKhldqgZXyCIyUPTVyMe7DoKk2MkoA3j8XXbwsizWPyrzUf', 'V8VPeorC', 1, NULL, NULL, '2024-06-25 08:19:10'),
('ncMtd2qReGO4P1rplH148phOhX7oTx6khdTZm0vd1ZSk7r3H5fMUVGSmUGeCMAVc', 'BLlXPWaZ', 1, NULL, NULL, '2024-06-25 14:33:16'),
('OtUevyBa8DyZa1BY5NES1JzWViin8B6v61Oue5mFHaDCE8nPJ3QzxwB3EeQbHaZO', 'uUv1D0V5', 1, NULL, NULL, '2024-06-24 18:31:03'),
('R22B18VofdQaq7UmWjNtjlVI2L9uGKzdzNDQnU1lho5s3Qf48PydpV4fYmglKCBb', 'iugZOWd7', 1, NULL, NULL, '2024-06-24 18:33:19'),
('rI8HzOKdToPLOrbEEHNO3FhdAKg8mduhH7AE49ngdIJTy0pZzE7ZZsNue6Qu6joD', 'm2gnq7iM', 1, NULL, NULL, '2024-06-24 18:23:41'),
('sjF0uzsYCQAWfesFHBNUfXF2JDdTJhj0GHch2i8yyw15KGYtIFpuhtYPUguXAj5w', 'HxTURLiQ', 1, NULL, NULL, '2024-06-24 18:31:31'),
('t44Zi2uz1z2pJ5kvg6LVjKFaheckgaUK8RhAbdF2P5ywxSSoK8cJgUfHTNiAE12k', 'NMOHRJD6', 1, NULL, NULL, '2024-06-24 18:25:00'),
('tiBNan6X71UN5pmwEvkpTIfWV8FY4M91NUyUsLGMhu4JcNxtCUBHrhDszcu60Lts', 'WzYmfsVp', 1, NULL, NULL, '2024-06-25 08:19:10'),
('v413Jd6aPDtJgwAAr6PPlwg59fBWK76svt2ExQDyqfLci76ffL4XK9B3UF5HWU6f', 'wQGKzxuC', 8, NULL, NULL, '2024-06-25 14:33:13'),
('vRc3bN0bp8a3KxdfnoSWD7WiQ0W2aFAmxGKxZjf6y8DFy0RTKkg4a5pGFkTIxwdl', 'C0qDaXgW', 1, NULL, NULL, '2024-06-24 18:33:53'),
('WHyFD5RWu6CXkhsH5IJnvsYgBF4kuiIwXb2t8jvkgWCzPsOj90A9tng5RzuvEc01', 'vKtMZraV', 1, NULL, NULL, '2024-06-24 18:37:13'),
('wW0G4HMxN8fJzA0JTX4JwLeIwOBFVyeevwAf6TP6HNLNxX0Q3dou0CKXPnITbMjB', 'eWgC6cyb', 1, NULL, NULL, '2024-06-25 08:19:11'),
('YbeW9ANnZX39EXAlG4UzqLEhpqrWZIPiAEGofFdkUat7ePAqMpbzJwTVgKTq5b6s', '7chRUwbZ', 1, NULL, NULL, '2024-06-24 18:25:00'),
('Zp5jGRH3oHKCj14HDNJ4x2MFBqqwoHgqkdGQ95fiPtLBSRi8yyOMHLtyUnFuCgbm', 'Rc7mix9d', 1, NULL, NULL, '2023-10-14 17:10:28'),
('ztVfprvcalzGWO3F9YJkCOZ21vcm97AYEjPlONJZZj7pj9e2mKtvyelTH1iLg77i', 'yJVJuQcX', 1, NULL, NULL, '2024-06-24 18:33:55');

-- --------------------------------------------------------

--
-- Table structure for table `ry_categories`
--

CREATE TABLE `ry_categories` (
  `royacateg_id` int(11) NOT NULL,
  `royacateg_category` varchar(255) DEFAULT NULL,
  `royacateg_icon` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_categories`
--

INSERT INTO `ry_categories` (`royacateg_id`, `royacateg_category`, `royacateg_icon`) VALUES
(1, 'Appartements', 'building'),
(2, 'house', 'House'),
(3, 'room', 'room'),
(4, 'office', 'Bureau'),
(5, 'villa', 'villa'),
(6, 'farm', 'farm'),
(7, 'riad', 'Tray');

-- --------------------------------------------------------

--
-- Table structure for table `ry_clients`
--

CREATE TABLE `ry_clients` (
  `ryclien_id` int(11) NOT NULL,
  `ryclien_user_id` int(11) NOT NULL,
  `ryclien_cin` varchar(255) DEFAULT NULL,
  `ryclien_annonceId` int(11) DEFAULT 0,
  `ryclien_prix` varchar(255) DEFAULT NULL,
  `ryclien_statue` varchar(20) DEFAULT NULL,
  `ryclien_is_blocked` int(11) DEFAULT 0,
  `ryclien_crdate` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_clients`
--

INSERT INTO `ry_clients` (`ryclien_id`, `ryclien_user_id`, `ryclien_cin`, `ryclien_annonceId`, `ryclien_prix`, `ryclien_statue`, `ryclien_is_blocked`, `ryclien_crdate`) VALUES
(1, 2, 'Jc121314', 1, NULL, NULL, 0, '2023-09-08 15:51:59'),
(2, 1, 'Jc131415', 2, NULL, NULL, 0, '2023-09-14 18:46:30'),
(3, 7, NULL, 3, NULL, NULL, 0, '2023-09-14 20:06:03'),
(4, 8, 'Jc131517', 4, NULL, NULL, 0, '2023-09-14 20:10:12'),
(5, 9, 'CC112233', 5, NULL, NULL, 0, '2023-09-14 20:12:08'),
(6, 10, NULL, 6, NULL, NULL, 0, '2023-09-22 08:47:50'),
(7, 10, NULL, 7, NULL, NULL, 0, '2023-09-23 10:50:05'),
(8, 10, NULL, 8, NULL, NULL, 0, '2023-09-23 10:55:21'),
(9, 11, NULL, 9, NULL, NULL, 0, '2023-09-23 12:15:35'),
(10, 10, NULL, 10, NULL, NULL, 0, '2023-09-23 12:44:13'),
(11, 12, NULL, 11, NULL, NULL, 0, '2023-09-23 13:03:07'),
(12, 13, NULL, 0, NULL, NULL, 0, '2023-10-06 18:52:34');

-- --------------------------------------------------------

--
-- Table structure for table `ry_device_info`
--

CREATE TABLE `ry_device_info` (
  `royadvcinf_id` bigint(20) NOT NULL,
  `royadvcinf_uuid` varchar(255) DEFAULT NULL,
  `royadvcinf_fcm_token` varchar(255) DEFAULT NULL,
  `royadvcinf_hardware` tinytext DEFAULT NULL,
  `royadvcinf_model` varchar(255) DEFAULT NULL,
  `royadvcinf_os` varchar(255) DEFAULT NULL,
  `royadvcinf_os_version` varchar(255) DEFAULT NULL,
  `royadvcinf_ip` varchar(15) DEFAULT NULL,
  `royadvcinf_crdate` timestamp NOT NULL DEFAULT current_timestamp(),
  `royadvcinf_userid` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_device_info`
--

INSERT INTO `ry_device_info` (`royadvcinf_id`, `royadvcinf_uuid`, `royadvcinf_fcm_token`, `royadvcinf_hardware`, `royadvcinf_model`, `royadvcinf_os`, `royadvcinf_os_version`, `royadvcinf_ip`, `royadvcinf_crdate`, `royadvcinf_userid`) VALUES
(6, '86610b1a-4a4e-5702-acae-2c49aa8227cb', 'ccUYWEMrQs6NnDGnuSjG9S:APA91bFuk1gQ0q1zCwhUvHdYIlNzlOGa_S-MQ2_Bb_l7VdqvDGITFQreo6TAwobb1mEhpzKNQB9_8I7-XWvuvp2YidLZq4xgKmfpBAhmaA_LMeTqzS4NtzxGZsvQSIdyf_owJR7Of3s2', 'Redmi', '2201117TY', 'Android', 'Android 12', '105.67.3.80', '2023-10-02 10:51:15', 1),
(7, '92012d97-f5b8-5fb7-a112-6beb30cfa8ef', 'djmQwR9VT9i4QTzxiQagzo:APA91bHrrgvAA4IL6rKBNqVkRJ6GIQfCkwlrkHQlyzMsSROdym9gr6cG8vfSYfvHMPSm0xq0oszdqbNDFeV-ScHydLq_vbxu16Uwuds_5spffN7obOqGywJxjBNAMvw11fZSb2_AzKCU', 'google', 'sdk_gphone_x86', 'Android', 'Android 11', '105.67.3.80', '2023-10-02 14:38:40', 3),
(8, '8c092b7d-3c13-5724-92a8-cc58bb005c3a', 'fxRxDJ23S7eGZX9HSTaOQJ:APA91bFeZ6IH7viLF3zHQdWVWy8IYvE3tqLLy9XOU7bafOdHQbdl2C-2K1YUlc9RKq1vUaDPQPSMENy-0Vn_jScHfSjUo2OJxvLXt2cMHyE0wrjAZEEVoTAqMmqB-z2oz2_t4co_TnYJ', 'samsung', 'SM-J710F', 'Android', 'Android 8.1.0', '196.200.176.254', '2024-06-24 18:23:37', 1);

-- --------------------------------------------------------

--
-- Table structure for table `ry_invoices`
--

CREATE TABLE `ry_invoices` (
  `ryinvoc_invoice_id` int(11) NOT NULL,
  `ryinvoc_client_id` int(11) DEFAULT NULL,
  `ryinvoc_annonce_id` int(11) DEFAULT NULL,
  `ryinvoc_invoice_date` datetime DEFAULT NULL,
  `ryinvoc_invoice_month` varchar(20) DEFAULT NULL,
  `ryinvoc_invoice_status` varchar(255) DEFAULT NULL,
  `ryinvoc_invoice_crdate` timestamp NULL DEFAULT current_timestamp(),
  `ryinvoc_invoice_prix` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `ry_invoices`
--

INSERT INTO `ry_invoices` (`ryinvoc_invoice_id`, `ryinvoc_client_id`, `ryinvoc_annonce_id`, `ryinvoc_invoice_date`, `ryinvoc_invoice_month`, `ryinvoc_invoice_status`, `ryinvoc_invoice_crdate`, `ryinvoc_invoice_prix`) VALUES
(1, 4, 4, '2023-09-23 16:41:35', 'September', 'rent', '2023-09-23 15:41:35', '6500'),
(1, 4, 4, '2023-09-23 16:41:35', 'September', 'rent', '2023-09-23 15:41:35', '6500');

-- --------------------------------------------------------

--
-- Table structure for table `ry_sms_verification`
--

CREATE TABLE `ry_sms_verification` (
  `ryusrvs_status` int(11) DEFAULT NULL,
  `ryusrvs_verification_code` varchar(20) DEFAULT NULL,
  `ryusrvs_mobile_number` varchar(20) DEFAULT NULL,
  `ryusrvs_crdate` timestamp NOT NULL DEFAULT current_timestamp(),
  `ryusrvs_verified_on` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_sms_verification`
--

INSERT INTO `ry_sms_verification` (`ryusrvs_status`, `ryusrvs_verification_code`, `ryusrvs_mobile_number`, `ryusrvs_crdate`, `ryusrvs_verified_on`) VALUES
(1, '544733', '0650421408', '2024-06-21 20:04:26', '2024-06-21 21:04:52'),
(1, '442893', '0650421408', '2024-06-22 22:18:14', '2024-06-22 22:18:26'),
(1, '362902', '0650421408', '2024-06-24 18:24:39', '2024-06-24 18:24:58'),
(NULL, '599751', '0650421408', '2024-06-24 18:42:39', NULL),
(1, '662320', '0650421408', '2024-06-24 18:43:16', '2024-06-24 18:43:34'),
(1, '968652', '0650421408', '2024-06-25 09:00:10', '2024-06-25 09:00:23'),
(1, '780873', '0650421408', '2024-06-25 14:32:29', '2024-06-25 14:32:46');

-- --------------------------------------------------------

--
-- Table structure for table `ry_user`
--

CREATE TABLE `ry_user` (
  `ryusr_id` int(11) NOT NULL,
  `ryusr_first_name` varchar(255) NOT NULL,
  `ryusr_last_name` varchar(255) NOT NULL,
  `ryusr_role` varchar(255) DEFAULT 'user',
  `ryusr_email` varchar(255) DEFAULT NULL,
  `ryusr_avatar` varchar(255) DEFAULT NULL,
  `ryusr_blocked` int(11) DEFAULT 0,
  `ryusr_mobile_number` varchar(20) DEFAULT NULL,
  `ryusr_discription` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_user`
--

INSERT INTO `ry_user` (`ryusr_id`, `ryusr_first_name`, `ryusr_last_name`, `ryusr_role`, `ryusr_email`, `ryusr_avatar`, `ryusr_blocked`, `ryusr_mobile_number`, `ryusr_discription`) VALUES
(1, 'ouknik', 'mhamdi', 'admin', 'ouknikabdeallah@gmai', '', 0, '0650421408', 'best smsar in taroudant city 🏙️'),
(2, 'abdellah d', 'ouknik', 'user', 'oopokni@gmail.com', '', 0, '0636420236', 'my father is a best smsar in taroudant city '),
(3, 'fatima', 'khti', 'user', NULL, NULL, 0, '0645447542', NULL),
(4, 'bsh', 'fr', 'user', NULL, NULL, 0, '065555555555', NULL),
(5, 'najat', 'najat', 'user', NULL, NULL, 0, '061212121212', NULL),
(6, 'said', 'said', 'user', NULL, NULL, 0, '0612131415', NULL),
(7, 'ahmed', 'lhamidi', 'user', NULL, NULL, 0, '06121314515', NULL),
(8, 'test', 'cliente', 'user', '', '', 0, '0611122111', ''),
(9, 'test 21', 'by client', 'user', 'test@gmail.com', '', 0, '0675149856', 'test by test'),
(10, 'Ahmed', 'Mhamdi', 'user', NULL, NULL, 0, '0624058920', NULL),
(11, 'New', 'Client', 'user', NULL, NULL, 0, '0624058921', NULL),
(12, 'ljhkj', 'khkj', 'user', NULL, NULL, 0, '0654856235', NULL),
(13, 'youssef', 'amzil', 'user', NULL, NULL, 0, '0644887730', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `ry_user_token`
--

CREATE TABLE `ry_user_token` (
  `ryustkn_id` int(11) NOT NULL,
  `ryustkn_usr_id` int(11) NOT NULL,
  `ryustkn_token` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ry_user_token`
--

INSERT INTO `ry_user_token` (`ryustkn_id`, `ryustkn_usr_id`, `ryustkn_token`) VALUES
(1, 1, 'EstZMxbmuO7WkabMAdw9OH4YgSfIY1Eo'),
(2, 1, 'ZMBQLmjc6ulEWXOgCpqn8M4wsKjsJrdk'),
(3, 1, 'NXYsuOL3QbVDbzbx8R1yc7m3bednXiIg'),
(4, 1, 'Na0Fc5pm31E26iOfI0t5cHDLUf2k43EZ'),
(5, 1, 'Y21XgYLcYWJ1ddTzkUgslqMg6oO9panW'),
(6, 1, 'rxNNXfCTloMVITXpfegdmnzzuaXA7Q7Y'),
(7, 1, 'tPJsmQV5zPK9xtTXw1LgkCl9LO3ZleVU'),
(8, 1, 'am16kBcCeuOfOtyQaqKcKGeNfCDyNlcN'),
(9, 1, 'FxTc4Eir9GEBmdL3L0L1pArIzimE8drn'),
(10, 2, 'D8uDWPGBwZCSHrC5CaMYVvLAmymda8ke'),
(11, 3, 'QnSFYFH5SvfkC5EXoMwiqDjz46Fm8c9v'),
(12, 4, 'lSRJwhmsUR0uZ00FJQKQgq4lOwxRB6iT'),
(13, 5, 'flNVNypSbYJJ66RrBw76MpS3UoVMXIYp'),
(14, 6, 'y6ekQhvZu8xfDyVjtE6o5JAa7UqQnSHL'),
(15, 7, 'VHg0uawkVrRaqHsLttPnC0PbLbaIS9pg'),
(16, 8, '0TFCKpx3jQ9JyqwsBRu6B4XMxwMV6YuW'),
(17, 9, 'oNhJ6mAm4cKSSNJlPIicW5LTCN9HrF5u'),
(18, 10, 'Xoy3y8hvmbJgPfXDHfgNOLtC8j8CSbRE'),
(19, 11, 'TVLoCGHc0mRizYH1TzlddG09peszJ2nr'),
(20, 12, 'KUr639z0LybW0ppxSnUE7o9ZymcVUQ5S'),
(21, 13, 'NymFNBY0wWLMUa5SolIZGHM9C34o7wIy');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `age` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `age`) VALUES
(1, 'Bob', 'bob@bob.com', 26);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `ry_abilities`
--
ALTER TABLE `ry_abilities`
  ADD PRIMARY KEY (`ryabilit_id`);

--
-- Indexes for table `ry_annonce`
--
ALTER TABLE `ry_annonce`
  ADD PRIMARY KEY (`royaannonse_id`);

--
-- Indexes for table `ry_annonce_abilities_avariable`
--
ALTER TABLE `ry_annonce_abilities_avariable`
  ADD PRIMARY KEY (`ryannoabilitvarb_id`);

--
-- Indexes for table `ry_annonce_images_avariable`
--
ALTER TABLE `ry_annonce_images_avariable`
  ADD PRIMARY KEY (`royaannimgava_id`);

--
-- Indexes for table `ry_annonce_message`
--
ALTER TABLE `ry_annonce_message`
  ADD PRIMARY KEY (`ryannmess_id`);

--
-- Indexes for table `ry_annose_images`
--
ALTER TABLE `ry_annose_images`
  ADD PRIMARY KEY (`royaimg_id`);

--
-- Indexes for table `ry_api_otts`
--
ALTER TABLE `ry_api_otts`
  ADD PRIMARY KEY (`royaapiot_id`);

--
-- Indexes for table `ry_categories`
--
ALTER TABLE `ry_categories`
  ADD PRIMARY KEY (`royacateg_id`);

--
-- Indexes for table `ry_clients`
--
ALTER TABLE `ry_clients`
  ADD PRIMARY KEY (`ryclien_id`);

--
-- Indexes for table `ry_device_info`
--
ALTER TABLE `ry_device_info`
  ADD PRIMARY KEY (`royadvcinf_id`);

--
-- Indexes for table `ry_user`
--
ALTER TABLE `ry_user`
  ADD PRIMARY KEY (`ryusr_id`);

--
-- Indexes for table `ry_user_token`
--
ALTER TABLE `ry_user_token`
  ADD PRIMARY KEY (`ryustkn_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `ry_annonce`
--
ALTER TABLE `ry_annonce`
  MODIFY `royaannonse_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `ry_annonce_abilities_avariable`
--
ALTER TABLE `ry_annonce_abilities_avariable`
  MODIFY `ryannoabilitvarb_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=91;

--
-- AUTO_INCREMENT for table `ry_annonce_images_avariable`
--
ALTER TABLE `ry_annonce_images_avariable`
  MODIFY `royaannimgava_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `ry_annonce_message`
--
ALTER TABLE `ry_annonce_message`
  MODIFY `ryannmess_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=47;

--
-- AUTO_INCREMENT for table `ry_annose_images`
--
ALTER TABLE `ry_annose_images`
  MODIFY `royaimg_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `ry_categories`
--
ALTER TABLE `ry_categories`
  MODIFY `royacateg_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `ry_clients`
--
ALTER TABLE `ry_clients`
  MODIFY `ryclien_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `ry_device_info`
--
ALTER TABLE `ry_device_info`
  MODIFY `royadvcinf_id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `ry_user`
--
ALTER TABLE `ry_user`
  MODIFY `ryusr_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `ry_user_token`
--
ALTER TABLE `ry_user_token`
  MODIFY `ryustkn_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
