# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: version.proto
# plugin: python-betterproto
# This file has been @generated

from dataclasses import dataclass

import betterproto


@dataclass(eq=False, repr=False)
class Version(betterproto.Message):
    """
    Currently ahnlich uses the major versions to interact with clients provided they match.
    """

    major: int = betterproto.uint32_field(1)
    minor: int = betterproto.uint32_field(2)
    patch: int = betterproto.uint32_field(3)
