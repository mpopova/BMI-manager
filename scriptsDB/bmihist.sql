-- phpMyAdmin SQL Dump
-- version 4.2.11
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Feb 19, 2016 at 07:16 PM
-- Server version: 5.6.21
-- PHP Version: 5.5.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


-- Database: `bmi`
--

-- --------------------------------------------------------

--
-- Table structure for table `bmihist`
--

CREATE TABLE IF NOT EXISTS `bmihist` (
`id` int(11) NOT NULL,
  `UserName` varchar(20) NOT NULL,
  `Date` date NOT NULL,
  `BMI` float NOT NULL
) 