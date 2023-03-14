<?php

namespace App\Controller;

use Doctrine\Persistence\ObjectManager;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\Form\FormInterface;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;

use App\Entity\User;
use App\Form\Type\SignUpType;

abstract class AuthBaseController extends AbstractController
{
    protected function __construct(
        protected ObjectManager $entityManager,
        protected UserPasswordHasherInterface $passwordHasher,
    ) {
    }

    protected function createSignUpForm(?User $user = null): FormInterface
    {
        return $this->createForm(SignUpType::class, $user, [
            'action' => $this->generateUrl('auth_sign_up_post'),
        ]);
    }

    protected function signUpUser(User $user): User
    {
        $hashedPassword = $this->passwordHasher->hashPassword(
            $user,
            $user->getPlainPassword()
        );
        $user->setPassword($hashedPassword);
        $user->setCreatedAt();
        $user->setUpdatedAt();
        $user->eraseCredentials();

        $this->entityManager->persist($user);
        $this->entityManager->flush();

        return $user;
    }
}
