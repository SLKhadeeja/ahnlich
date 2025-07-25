# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: db/query.proto
# plugin: python-betterproto
# This file has been @generated

from dataclasses import dataclass
from typing import List

import betterproto

from ... import keyval as __keyval__
from ... import predicates as __predicates__
from ...algorithm import algorithms as __algorithm_algorithms__
from ...algorithm import nonlinear as __algorithm_nonlinear__


@dataclass(eq=False, repr=False)
class CreateStore(betterproto.Message):
    """
    Creates a new store in the database with the specified dimension, predicates, and non-linear indices.
     If `error_if_exists` is set to true, it will return an error if the store already exists.
    """

    store: str = betterproto.string_field(1)
    dimension: int = betterproto.uint32_field(2)
    create_predicates: List[str] = betterproto.string_field(3)
    non_linear_indices: List["__algorithm_nonlinear__.NonLinearAlgorithm"] = (
        betterproto.enum_field(4)
    )
    error_if_exists: bool = betterproto.bool_field(5)


@dataclass(eq=False, repr=False)
class GetKey(betterproto.Message):
    """Retrieves values from the store based on provided keys."""

    store: str = betterproto.string_field(1)
    keys: List["__keyval__.StoreKey"] = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class GetPred(betterproto.Message):
    """
    Retrieves values from the store based on predicates. Validation checks if the predicate is enabled.
    """

    store: str = betterproto.string_field(1)
    condition: "__predicates__.PredicateCondition" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class GetSimN(betterproto.Message):
    """
    Retrieves the `n` most similar items to the input vector from the store, using the specified algorithm.
     Validation checks that the dimensions of the input vector match the store's dimensions.
     `n` could be less than originally specified.
    """

    store: str = betterproto.string_field(1)
    search_input: "__keyval__.StoreKey" = betterproto.message_field(2)
    closest_n: int = betterproto.uint64_field(3)
    algorithm: "__algorithm_algorithms__.Algorithm" = betterproto.enum_field(4)
    condition: "__predicates__.PredicateCondition" = betterproto.message_field(5)


@dataclass(eq=False, repr=False)
class CreatePredIndex(betterproto.Message):
    """
    Creates an index in the store based on the provided predicates.
     This operation is idempotent: it will only add new predicates, not remove existing ones.
    """

    store: str = betterproto.string_field(1)
    predicates: List[str] = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class CreateNonLinearAlgorithmIndex(betterproto.Message):
    """
    Creates an index for non-linear algorithms in the store, if it does not already exist.
     This operation is idempotent: it will only add new non-linear indices, not remove existing ones.
    """

    store: str = betterproto.string_field(1)
    non_linear_indices: List["__algorithm_nonlinear__.NonLinearAlgorithm"] = (
        betterproto.enum_field(2)
    )


@dataclass(eq=False, repr=False)
class DropPredIndex(betterproto.Message):
    """
    Drops the specified predicates from the store.
     If `error_if_not_exists` is true, an error is returned if the predicate does not exist.
    """

    store: str = betterproto.string_field(1)
    predicates: List[str] = betterproto.string_field(2)
    error_if_not_exists: bool = betterproto.bool_field(3)


@dataclass(eq=False, repr=False)
class DropNonLinearAlgorithmIndex(betterproto.Message):
    """
    Drops the specified non-linear indices from the store.
     If `error_if_not_exists` is true, an error is returned if the non-linear index does not exist.
    """

    store: str = betterproto.string_field(1)
    non_linear_indices: List["__algorithm_nonlinear__.NonLinearAlgorithm"] = (
        betterproto.enum_field(2)
    )
    error_if_not_exists: bool = betterproto.bool_field(3)


@dataclass(eq=False, repr=False)
class DelKey(betterproto.Message):
    """
    Deletes the specified keys from the store and returns the number of deleted keys.
     It will also update the indices in a non-blocking manner.
    """

    store: str = betterproto.string_field(1)
    keys: List["__keyval__.StoreKey"] = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class DelPred(betterproto.Message):
    """
    Deletes values from the store based on the provided predicate condition.
     It will also update the indices in a non-blocking manner.
    """

    store: str = betterproto.string_field(1)
    condition: "__predicates__.PredicateCondition" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class DropStore(betterproto.Message):
    """
    Drops a store and deletes all its data and associated indices.
     If `error_if_not_exists` is true, it will return an error if the store does not exist.
    """

    store: str = betterproto.string_field(1)
    error_if_not_exists: bool = betterproto.bool_field(2)


@dataclass(eq=False, repr=False)
class InfoServer(betterproto.Message):
    """A request to get server information such as host, port, and version."""

    pass


@dataclass(eq=False, repr=False)
class ListStores(betterproto.Message):
    """
    A request to list all the stores on the server, along with their size or length.
    """

    pass


@dataclass(eq=False, repr=False)
class ListClients(betterproto.Message):
    """A request to list all the clients currently connected to the server."""

    pass


@dataclass(eq=False, repr=False)
class Ping(betterproto.Message):
    """A simple ping request to check if the server is reachable."""

    pass


@dataclass(eq=False, repr=False)
class Set(betterproto.Message):
    """
    A request to set multiple key-value entries in the store.
     Validation is done for each vector before updating the store.
    """

    store: str = betterproto.string_field(1)
    inputs: List["__keyval__.DbStoreEntry"] = betterproto.message_field(2)
