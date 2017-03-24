-- phpMyAdmin SQL Dump
-- version 4.4.13.1deb1
-- http://www.phpmyadmin.net
--
-- Client :  localhost
-- Généré le :  Mar 23 Février 2016 à 15:24
-- Version du serveur :  5.6.28-0ubuntu0.15.10.1
-- Version de PHP :  5.6.11-1ubuntu3.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données :  `db_azzinoth`
--

-- --------------------------------------------------------

--
-- Structure de la table `clients`
--

CREATE TABLE IF NOT EXISTS `clients` (
  `reference` varchar(50) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `configurations`
--

CREATE TABLE IF NOT EXISTS `configurations` (
  `id` int(11) NOT NULL,
  `client_reference` varchar(50) NOT NULL,
  `path` varchar(255) NOT NULL,
  `stream_name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `historized_need`
--

CREATE TABLE IF NOT EXISTS `historized_need` (
  `historized_id` int(11) NOT NULL,
  `need_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `historized_timetable`
--

CREATE TABLE IF NOT EXISTS `historized_timetable` (
  `id` int(11) NOT NULL,
  `regular_id` int(11) NOT NULL,
  `process_id` int(11) DEFAULT NULL,
  `scheduled_date` date NOT NULL,
  `minimal_start_time` time DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `maximal_start_time` time DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `status` varchar(50) DEFAULT 'scheduled',
  `infos` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `migrations`
--

CREATE TABLE IF NOT EXISTS `migrations` (
  `migration` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Structure de la table `regular_days`
--

CREATE TABLE IF NOT EXISTS `regular_days` (
  `regular_id` int(11) NOT NULL,
  `day` enum('monday','tuesday','wednesday','thursday','friday','saturday','sunday') NOT NULL,
  `minimal_start_time` time DEFAULT NULL,
  `maximal_start_time` time DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `regular_need`
--

CREATE TABLE IF NOT EXISTS `regular_need` (
  `regular_id` int(11) NOT NULL,
  `need_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `regular_timetable`
--

CREATE TABLE IF NOT EXISTS `regular_timetable` (
  `id` int(11) NOT NULL,
  `service_reference` varchar(50) NOT NULL,
  `config_id` int(11) NOT NULL,
  `process_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `scheduled_need`
--

CREATE TABLE IF NOT EXISTS `scheduled_need` (
  `scheduled_id` int(11) NOT NULL,
  `need_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `scheduled_timetable`
--

CREATE TABLE IF NOT EXISTS `scheduled_timetable` (
  `id` int(11) NOT NULL,
  `regular_id` int(11) NOT NULL,
  `process_id` int(11) DEFAULT NULL,
  `scheduled_date` date NOT NULL,
  `minimal_start_time` time DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `maximal_start_time` time DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `status` varchar(50) DEFAULT 'scheduled',
  `infos` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `servers`
--

CREATE TABLE IF NOT EXISTS `servers` (
  `reference` varchar(50) NOT NULL,
  `dns` varchar(255) NOT NULL DEFAULT '.kdata.fr',
  `type` enum('calc','data') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `services`
--

CREATE TABLE IF NOT EXISTS `services` (
  `reference` varchar(50) NOT NULL,
  `service_reference` varchar(50) NOT NULL,
  `calc_server_reference` varchar(50) NOT NULL,
  `last_activity` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `services_types`
--

CREATE TABLE IF NOT EXISTS `services_types` (
  `reference` varchar(50) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `status_types`
--

CREATE TABLE IF NOT EXISTS `status_types` (
  `reference` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Index pour les tables exportées
--

--
-- Index pour la table `clients`
--
ALTER TABLE `clients`
  ADD PRIMARY KEY (`reference`);

--
-- Index pour la table `configurations`
--
ALTER TABLE `configurations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `configurations_fk_client_reference` (`client_reference`);

--
-- Index pour la table `historized_need`
--
ALTER TABLE `historized_need`
  ADD PRIMARY KEY (`historized_id`,`need_id`),
  ADD KEY `historized_need_fk_need_id` (`need_id`);

--
-- Index pour la table `historized_timetable`
--
ALTER TABLE `historized_timetable`
  ADD PRIMARY KEY (`id`),
  ADD KEY `historized_timetable_fk_regular_id` (`regular_id`),
  ADD KEY `historized_timetable_fk_status_reference` (`status`);

--
-- Index pour la table `regular_days`
--
ALTER TABLE `regular_days`
  ADD PRIMARY KEY (`regular_id`,`day`);

--
-- Index pour la table `regular_need`
--
ALTER TABLE `regular_need`
  ADD PRIMARY KEY (`regular_id`,`need_id`),
  ADD KEY `regular_need_fk_need_id` (`need_id`);

--
-- Index pour la table `regular_timetable`
--
ALTER TABLE `regular_timetable`
  ADD PRIMARY KEY (`id`),
  ADD KEY `regular_timetable_fk_service_reference` (`service_reference`),
  ADD KEY `regular_timetable_fk_config_id` (`config_id`);

--
-- Index pour la table `scheduled_need`
--
ALTER TABLE `scheduled_need`
  ADD PRIMARY KEY (`scheduled_id`,`need_id`),
  ADD KEY `scheduled_need_fk_need_id` (`need_id`);

--
-- Index pour la table `scheduled_timetable`
--
ALTER TABLE `scheduled_timetable`
  ADD PRIMARY KEY (`id`),
  ADD KEY `scheduled_timetable_fk_regular_id` (`regular_id`),
  ADD KEY `scheduled_timetable_fk_status_reference` (`status`);

--
-- Index pour la table `servers`
--
ALTER TABLE `servers`
  ADD PRIMARY KEY (`reference`);

--
-- Index pour la table `services`
--
ALTER TABLE `services`
  ADD PRIMARY KEY (`reference`),
  ADD KEY `services_fk_service_reference` (`service_reference`),
  ADD KEY `services_fk_calc_server_reference` (`calc_server_reference`);

--
-- Index pour la table `services_types`
--
ALTER TABLE `services_types`
  ADD PRIMARY KEY (`reference`);

--
-- Index pour la table `status_types`
--
ALTER TABLE `status_types`
  ADD PRIMARY KEY (`reference`);

--
-- AUTO_INCREMENT pour les tables exportées
--

--
-- AUTO_INCREMENT pour la table `configurations`
--
ALTER TABLE `configurations`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT pour la table `historized_timetable`
--
ALTER TABLE `historized_timetable`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT pour la table `regular_timetable`
--
ALTER TABLE `regular_timetable`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT pour la table `scheduled_timetable`
--
ALTER TABLE `scheduled_timetable`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- Contraintes pour les tables exportées
--

--
-- Contraintes pour la table `configurations`
--
ALTER TABLE `configurations`
  ADD CONSTRAINT `configurations_fk_client_reference` FOREIGN KEY (`client_reference`) REFERENCES `clients` (`reference`);

--
-- Contraintes pour la table `historized_need`
--
ALTER TABLE `historized_need`
  ADD CONSTRAINT `historized_need_fk_historized_id` FOREIGN KEY (`historized_id`) REFERENCES `historized_timetable` (`id`),
  ADD CONSTRAINT `historized_need_fk_need_id` FOREIGN KEY (`need_id`) REFERENCES `historized_timetable` (`id`);

--
-- Contraintes pour la table `historized_timetable`
--
ALTER TABLE `historized_timetable`
  ADD CONSTRAINT `historized_timetable_fk_regular_id` FOREIGN KEY (`regular_id`) REFERENCES `regular_timetable` (`id`),
  ADD CONSTRAINT `historized_timetable_fk_status_reference` FOREIGN KEY (`status`) REFERENCES `status_types` (`reference`);

--
-- Contraintes pour la table `regular_days`
--
ALTER TABLE `regular_days`
  ADD CONSTRAINT `regular_weeks_fk_regular_id` FOREIGN KEY (`regular_id`) REFERENCES `regular_timetable` (`id`);

--
-- Contraintes pour la table `regular_need`
--
ALTER TABLE `regular_need`
  ADD CONSTRAINT `regular_need_fk_need_id` FOREIGN KEY (`need_id`) REFERENCES `regular_timetable` (`id`),
  ADD CONSTRAINT `regular_need_fk_regular_id` FOREIGN KEY (`regular_id`) REFERENCES `regular_timetable` (`id`);

--
-- Contraintes pour la table `regular_timetable`
--
ALTER TABLE `regular_timetable`
  ADD CONSTRAINT `regular_timetable_fk_config_id` FOREIGN KEY (`config_id`) REFERENCES `configurations` (`id`),
  ADD CONSTRAINT `regular_timetable_fk_service_reference` FOREIGN KEY (`service_reference`) REFERENCES `services` (`reference`);

--
-- Contraintes pour la table `scheduled_need`
--
ALTER TABLE `scheduled_need`
  ADD CONSTRAINT `scheduled_need_fk_need_id` FOREIGN KEY (`need_id`) REFERENCES `scheduled_timetable` (`id`),
  ADD CONSTRAINT `scheduled_need_fk_scheduled_id` FOREIGN KEY (`scheduled_id`) REFERENCES `scheduled_timetable` (`id`);

--
-- Contraintes pour la table `scheduled_timetable`
--
ALTER TABLE `scheduled_timetable`
  ADD CONSTRAINT `scheduled_timetable_fk_regular_id` FOREIGN KEY (`regular_id`) REFERENCES `regular_timetable` (`id`),
  ADD CONSTRAINT `scheduled_timetable_fk_status_reference` FOREIGN KEY (`status`) REFERENCES `status_types` (`reference`);

--
-- Contraintes pour la table `services`
--
ALTER TABLE `services`
  ADD CONSTRAINT `services_fk_calc_server_reference` FOREIGN KEY (`calc_server_reference`) REFERENCES `servers` (`reference`),
  ADD CONSTRAINT `services_fk_service_reference` FOREIGN KEY (`service_reference`) REFERENCES `services_types` (`reference`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
