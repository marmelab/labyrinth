<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Bundle\SecurityBundle\Security;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\Security\Http\Authentication\AuthenticationUtils;

#[Route('/auth')]
class SecurityController extends AbstractController
{
    #[Route(path: '/sign-in', name: 'app_login')]
    public function login(AuthenticationUtils $authenticationUtils): Response
    {
        if ($this->getUser()) {
            return $this->redirectToRoute('home');
        }

        $error = $authenticationUtils->getLastAuthenticationError();
        $lastUsername = $authenticationUtils->getLastUsername();
        return $this->render('auth/sign_in.html.twig', ['last_username' => $lastUsername, 'error' => $error]);
    }

    #[Route(path: '/sign-out', name: 'auth_sign_out')]
    public function logout(Security $security): Response
    {
        return $security->logout();
    }
}
